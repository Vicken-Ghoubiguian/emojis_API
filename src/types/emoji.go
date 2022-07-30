package types

// Defining the type 'Emoji' which define an emoji...
type Emoji struct {
	number          int
	title           string
	unicode         string
	cldr_short_name string
	category        Category
	subcategory     SubCategory
}

//
func (current_emoji *Emoji) GetNumber() int {

	return current_emoji.number
}

//
func (current_emoji *Emoji) GetTitle() string {

	return current_emoji.title
}

//
func (current_emoji *Emoji) GetUnicode() string {

	return current_emoji.unicode
}

//
func (current_emoji *Emoji) GetCldr_short_name() string {

	return current_emoji.cldr_short_name
}

//
func (current_emoji *Emoji) GetCategory() Category {

	return current_emoji.category
}

//
func (current_emoji *Emoji) GetSubcategory() SubCategory {

	return current_emoji.subcategory
}
