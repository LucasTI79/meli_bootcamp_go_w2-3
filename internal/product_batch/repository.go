package productbatch

const (
	SaveQuery         = "INSERT INTO product_batches ( batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?,?,?,?,?,?,?,?,?,?)"
	ProductExists     = "SELECT id FROM products WHERE id=?"
	SectionExists     = "SELECT id FROM sections WHERE id=?"
	ProductsBySection = "SELECT count(pb.id) as `products_count`,pb.section_id, s.section_number FROM product_batches pb JOIN section ON pb.section_id = s.id GROUP BY pb.section_id"
)
