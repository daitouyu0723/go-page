package models

type page struct {
	PageLimit   int
	PageTotal   int
	CurrentPage int
	PrevPage    int //empty means disabled
	NextPage    int //empty means disabled
	ItemTotal   int
	ItemFirst   int
	ItemLast    int
	ItemOffset  int
}

func NewPage(currentPage int, pageLimit int, itemTotal int) *page {
	if pageLimit < 1 {
		pageLimit = 10
	}
	if currentPage < 1 {
		currentPage = 1
	}
	if itemTotal < 0 {
		itemTotal = 0
	}
	if itemTotal < pageLimit {
		return &page{
			PageLimit:   pageLimit,
			PageTotal:   1,
			CurrentPage: 1,
			PrevPage:    0,
			NextPage:    1,
			ItemTotal:   itemTotal,
			ItemFirst:   1,
			ItemLast:    itemTotal,
			ItemOffset:  0,
		}
	}
	p := &page{}
	p.PageLimit = pageLimit
	if itemTotal%pageLimit == 0 {
		p.PageTotal = itemTotal / pageLimit
	} else {
		p.PageTotal = itemTotal/pageLimit + 1
	}
	if currentPage < p.PageTotal {
		p.CurrentPage = currentPage
	} else {
		p.CurrentPage = p.PageTotal
	}
	p.PrevPage = p.CurrentPage - 1
	if p.CurrentPage == p.PageTotal {
		p.NextPage = p.PageTotal
	} else {
		p.NextPage = p.CurrentPage + 1
	}
	p.ItemTotal = itemTotal
	p.ItemFirst = (p.CurrentPage-1)*p.PageLimit + 1
	p.ItemLast = p.ItemFirst + p.PageLimit - 1
	if p.ItemLast > p.ItemTotal {
		p.ItemLast = p.ItemTotal
	}
	p.ItemOffset = p.ItemFirst - 1
	return p
}
