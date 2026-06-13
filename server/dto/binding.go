package dto

// binding 内置的验证包括：
// 不能为空，并且不能没有这个字段
// required：必填字段，如：binding:"required"

// 针对字符串的长度
// min 最小长度，如：binding:"min=5"
// max 最大长度，如：binding:"max=10"
// len 长度，如：binding:"len=6"

// 针对数字的大小
// eq 等于，如：binding:"eq=3"
// ne 不等于，如：binding:"ne=12"
// gt 大于，如：binding:"gt=10"
// gte 大于等于，如：binding:"gte=10"
// lt 小于，如：binding:"lt=10"
// lte 小于等于，如：binding:"lte=10"

// 针对同级字段的
// eqfield 等于其他字段的值，如：PassWord string `binding:"eqfield=Password"`
// nefield 不等于其他字段的值，如：PassWord string `binding:"nefield=Password"`

// - 忽略字段，如：binding:"-"

// 枚举，如：oneof=red green

// 字符串
// contains=abc  // 包含abc的字符串
// excludes // 不包含
// startswith  // 字符串前缀
// endswith  // 字符串后缀

// 数组
// dive 后面的验证就是针对数组中的每一个元素

// 网络验证
// ip
// ipv4
// ipv6
// uri
// url

// 日期验证
// datetime=2006-01-02
