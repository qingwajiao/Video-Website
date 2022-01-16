package dbops

import (
	"database/sql"
	"log"
)

type VideoHelper struct {
	Conn *sql.DB
}

func NewVideoHelper() *VideoHelper {
	return &VideoHelper{Conn: dbConn}
}

func (vh *VideoHelper) AddVideoDeletionRecord(vid string) error {
	stmtIns, err := vh.Conn.Prepare("INSERT INTO video_del_rec (video_id) VALUES(?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord error: %v", err)
		return err
	}

	defer stmtIns.Close()
	return nil
}

func (vh *VideoHelper) ReadVideoDeletionRecord(count int) ([]string, error) {
	stmtOut, err := vh.Conn.Prepare("SELECT video_id FROM video_del_rec LIMIT ?")

	var ids []string

	if err != nil {
		return ids, err
	}

	rows, err := stmtOut.Query(count)
	if err != nil {
		log.Printf("Query VideoDeletionRecord error: %v", err)
		return ids, err
	}

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	defer stmtOut.Close()
	return ids, nil
}

func (vh *VideoHelper) DelVideoDeletionRecord(vid string) error {
	stmtDel, err := vh.Conn.Prepare("DELETE FROM video_del_rec WHERE video_id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("Deleting VideoDeletionRecord error: %v", err)
		return err
	}

	defer stmtDel.Close()
	return nil
}
