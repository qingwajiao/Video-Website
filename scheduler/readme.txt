scheduler 调度器

什么时候需要用scheduler？
1、一般需要异步执行的任务
2、周期性的任务
如：本项目中删除视频的时候需要 延迟删除的需求，

API -> remove video -> insert videoID to video_clear_tables
start scheduler  VideoClearDispatcher get the videoID from video_clear_tables
VideoClearExecutor remove video from video_clear_tables





