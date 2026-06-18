package api

import "Helios/common"

// 分页信息
type PaginationConfig struct {
	PageNumber uint  `json:"page_number"` // 页码
	PageSize   uint  `json:"page_size"`   // 每页数据量
	Total      int64 `json:"total"`       // 数据量
	Paginable  bool  `json:"paginable"`   // 是否需要分页，默认 false，则不需要分页
}

// GetPaginationLimitAndOffset 计算分页查询的 limit 和 offset。
// limit=0, offset=0 表示不需要分页。
// 注意：不会修改原始 PaginationConfig 的字段值。
func (p *PaginationConfig) GetPaginationLimitAndOffset() (limit, offset int) {
	if !p.Paginable {
		return 0, 0
	}

	pageSize := p.PageSize
	if pageSize == 0 || pageSize > common.SystemPageSettingConfig.MaxPageSize {
		pageSize = common.SystemPageSettingConfig.PageSize
	}

	pageNumber := p.PageNumber
	if pageNumber == 0 {
		pageNumber = 1
	}

	// 使用 int64 做中间计算，防止 uint 乘法溢出
	offset64 := int64(pageNumber-1) * int64(pageSize)
	return int(pageSize), int(offset64)
}

// 数据分页返回格式
type PaginationResponse struct {
	Pagination PaginationConfig `json:"pagination"`
	List       any              `json:"list"`
}
