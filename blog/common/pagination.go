package common

func Pagination(count int, page int, page_size int) map[string]interface{} {
	total_page := 0
	if count % page_size == 0 {
		total_page = count / page_size
	} else {
		total_page = count / page_size + 1
	}
	pagination := map[string]interface{}{
		"total_post": count,
		"total_page": total_page,
		"has_next": page + 1 < total_page,
		"has_pre": page != 0,
	}
	return pagination
}
