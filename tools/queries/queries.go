package queries

const (
	GET_ALL_FOOD   = `SELECT * FROM tb_food WHERE food_status = 1`
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
						food_desc = ?,
					WHERE food_id = ? AND food_status = 1`
	DELETE_FOOD = `UPDATE tb_food
					SET food_status = 0
					WHERE food_id = ?`
)
