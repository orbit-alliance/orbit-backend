package gateway_42

import "time"

type ImageUserResponse struct {
	Link     string `json:"link"`
	Versions struct {
		Large  string `json:"large"`
		Medium string `json:"medium"`
		Small  string `json:"small"`
		Micro  string `json:"micro"`
	} `json:"versions"`
}

type ProjectsUsersResponse struct {
	ID            int    `json:"id"`
	Occurrence    int    `json:"occurrence"`
	FinalMark     int    `json:"final_mark"`
	Status        string `json:"status"`
	Validated     bool   `json:"validated?"`
	CurrentTeamID int    `json:"current_team_id"`
	Project       struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Slug     string `json:"slug"`
		ParentID any    `json:"parent_id"`
	} `json:"project"`
	CursusIds   []int     `json:"cursus_ids"`
	MarkedAt    time.Time `json:"marked_at"`
	Marked      bool      `json:"marked"`
	RetriableAt time.Time `json:"retriable_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AchievementsUserResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Tier         string `json:"tier"`
	Kind         string `json:"kind"`
	Visible      bool   `json:"visible"`
	Image        string `json:"image"`
	NbrOfSuccess any    `json:"nbr_of_success"`
	UsersURL     string `json:"users_url"`
}

type UserResponse struct {
	ID              int                        `json:"id"`
	Email           string                     `json:"email"`
	Login           string                     `json:"login"`
	FirstName       string                     `json:"first_name"`
	LastName        string                     `json:"last_name"`
	UsualFullName   string                     `json:"usual_full_name"`
	UsualFirstName  any                        `json:"usual_first_name"`
	URL             string                     `json:"url"`
	Phone           string                     `json:"phone"`
	Displayname     string                     `json:"displayname"`
	Kind            string                     `json:"kind"`
	Image           ImageUserResponse          `json:"image"`
	Staff           bool                       `json:"staff?"`
	CorrectionPoint int                        `json:"correction_point"`
	PoolMonth       string                     `json:"pool_month"`
	PoolYear        string                     `json:"pool_year"`
	Location        string                     `json:"location"`
	Wallet          int                        `json:"wallet"`
	AnonymizeDate   string                     `json:"anonymize_date"`
	DataErasureDate string                     `json:"data_erasure_date"`
	CreatedAt       time.Time                  `json:"created_at"`
	UpdatedAt       time.Time                  `json:"updated_at"`
	AlumnizedAt     any                        `json:"alumnized_at"`
	Alumni          bool                       `json:"alumni?"`
	Active          bool                       `json:"active?"`
	Groups          []any                      `json:"groups"`
	CursusUsers     []any                      `json:"cursus_users"`
	ProjectsUsers   []ProjectsUsersResponse    `json:"projects_users"`
	Achievements    []AchievementsUserResponse `json:"achievements"`
	Titles          []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"titles"`
	TitlesUsers []struct {
		ID        int       `json:"id"`
		UserID    int       `json:"user_id"`
		TitleID   int       `json:"title_id"`
		Selected  bool      `json:"selected"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"titles_users"`
	Partnerships    []any `json:"partnerships"`
	Patroned        []any `json:"patroned"`
	Patroning       []any `json:"patroning"`
	ExpertisesUsers []struct {
		ID          int       `json:"id"`
		ExpertiseID int       `json:"expertise_id"`
		Interested  bool      `json:"interested"`
		Value       int       `json:"value"`
		ContactMe   bool      `json:"contact_me"`
		CreatedAt   time.Time `json:"created_at"`
		UserID      int       `json:"user_id"`
	} `json:"expertises_users"`
	Roles       []any `json:"roles"`
	Campus      []any `json:"campus"`
	CampusUsers []struct {
		ID        int       `json:"id"`
		UserID    int       `json:"user_id"`
		CampusID  int       `json:"campus_id"`
		IsPrimary bool      `json:"is_primary"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"campus_users"`
}
