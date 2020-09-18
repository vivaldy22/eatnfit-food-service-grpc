package queries

const (
	GET_ALL_FOOD = `SELECT * 
					FROM tb_food
					WHERE food_status = 1 AND
					food_name LIKE ?
					ORDER BY 3
					LIMIT %v, %v`
	GET_TOTAL_FOOD = `SELECT COUNT(*) FROM tb_food WHERE food_status = 1`
	GET_BY_ID_FOOD = `SELECT * FROM tb_food WHERE food_id = ? AND food_status = 1`
	CREATE_FOOD    = `INSERT INTO tb_food VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, 1)`
	UPDATE_FOOD    = `UPDATE tb_food
					SET food_portion = ?,
						food_name = ?,
						food_calories = ?,
						food_fat = ?,
						food_carbs = ?,
						food_protein = ?,
						food_price = ?,
						food_desc = ?
					WHERE food_id = ? AND food_status = 1`
	DELETE_FOOD = `UPDATE tb_food
					SET food_status = 0
					WHERE food_id = ?`

	GET_ALL_PACKET = `SELECT * 
						FROM tb_packet
						WHERE packet_status = 1 AND
						packet_name LIKE ?
						ORDER BY 2
						LIMIT %v, %v`
	GET_TOTAL_PACKET       = `SELECT COUNT(*) FROM tb_packet WHERE packet_status = 1`
	GET_PACKET_BY_ID       = `SELECT * FROM tb_packet WHERE packet_id = ? AND packet_status = 1`
	GET_FOODS_BY_PACKET_ID = `SELECT f.*
								FROM tb_packet_and_food pf
								JOIN tb_food f ON pf.food_id = f.food_id
								JOIN tb_packet p ON pf.packet_id = p.packet_id
								WHERE p.packet_id = ?`
	CREATE_PACKET        = `INSERT INTO tb_packet VALUES (?, ?, ?, ?, ?, 1)`
	CREATE_DETAIL_PACKET = `INSERT INTO tb_packet_and_food VALUES (?, ?, ?, 1)`

	DELETE_PACKET = `UPDATE tb_packet
					SET packet_status = 0
					WHERE packet_id = ?`
	DELETE_DETAIL_PACKET = `UPDATE tb_packet_and_food
							SET pm_status = 0
							WHERE packet_id = ?`
	DELETE_PERMANENT_DETAIL_PACKET = `DELETE FROM tb_packet_and_food WHERE packet_id = ?`
	UPDATE_PACKET                  = `UPDATE tb_packet
					SET packet_name = ?,
						packet_portion = ?,
						packet_price = ?,
						packet_desc = ?
					WHERE packet_id = ?`
	UPDATE_DETAIL_PACKET = `UPDATE tb_packet_and_food
					SET food_id = ?
					WHERE packet_id = ?`

	GET_ALL_TRANSACTION = `SELECT * 
						FROM tb_transaction
						WHERE trans_status = 1
						ORDER BY 2
						LIMIT %v, %v`
	GET_TOTAL_TRANSACTION       = `SELECT COUNT(*) FROM tb_transaction WHERE trans_status = 1`
	GET_TRANS_BY_ID_TRANSACTION = `SELECT * FROM tb_transaction WHERE trans_id = ? AND trans_status = 1`
	GET_TRANS_BY_ID_USER        = `SELECT * FROM tb_transaction WHERE user_id = ? AND trans_status = 1`
	CREATE_TRANS                = `INSERT INTO tb_transaction
									VALUES (?, ?, ?, ?, ?, (SELECT packet_price FROM tb_packet WHERE packet_id = ?) * ? , 
									?, ?, ?, ?, 0, 1)`
	DELETE_TRANS = `UPDATE tb_transaction
									SET trans_status = 0
									WHERE trans_id = ?`
)
