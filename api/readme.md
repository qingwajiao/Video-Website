## API 设计：用户
# 创建 （注册）用户：URL:/user Method:POST,SC:201,400,500
# 用户登录：URL:/user/:username Method:POST,SC:200,400,500
# 获取用户基本信息：URL:/user/:username Method:GET,SC:200,400,403,500
# 用户注销：URL:/user/:username Method:DELETE,SC:204,400,401,403,500
