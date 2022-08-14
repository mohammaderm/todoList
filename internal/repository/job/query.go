package job

var (
	limit     = 3
	CreateJob = "insert into job (name, description, accountid) values ($1,$2,$3);"
	GetAllJob = "select * from job where accountid = $1 order by created_at desc limit $2 offset $3;"
	DeleteJob = "delete from job where id = $1 and accountid = $2;"
	UpdateJob = "update job set status = True where id = $1 and accountid = $2;"
)
