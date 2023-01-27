package models

type Follow struct {
	Follower      int          `json:"follower"`
	Following     int          `json:"follwoing"`
	UserFollower  UserResponse ` gorm:"foreignkey:Follower" json:"-"`
	UserFollowing UserResponse ` gorm:"foreignkey:Following" json:"-"`
}
