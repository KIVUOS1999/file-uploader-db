package store

const (
	INSERT_TO_FILE_DETAILS = `
		insert into file_details 
		(file_id,file_name,file_size,total_chunks,created_at)
		values 
		($1, $2, $3, $4, $5)
	`

	INSERT_TO_CHUNK_DETAILS = `
		insert into chunk_details
		(chunk_id,file_id,check_sum,chunk_order,created_at)
		values
		($1,$2,$3,$4,$5)
	`

	FETCH_FILE_DETAILS_BY_FILE_ID = `
		select * from file_details
		where
		file_id=$1
	`

	FETCH_CHUNK_BY_ORDER = `
		select chunk_id,check_sum,chunk_order from chunk_details
		where file_id=$1
		order by chunk_order asc
	`

	DELETE_CHUNK_BY_FILE_ID = `
		delete from chunk_details 
		where file_id=$1
	`

	DELETE_FILE_BY_FILE_ID = `
		delete from file_details
		where file_id=$1
	`

	// currently there is no user but in future there will be
	FETCH_FILE_PER_USER = `
	select file_id, file_name, file_size, created_at
	from file_details`
)
