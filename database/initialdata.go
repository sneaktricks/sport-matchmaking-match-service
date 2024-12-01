package database

// func match1() *model.Match {
// 	minParticipants := int32(2)
// 	maxParticipants := int32(4)
// 	createData := model.MatchCreate{
// 		Sport:             "Tennis",
// 		MinParticipants:   &minParticipants,
// 		MaxParticipants:   &maxParticipants,
// 		Location:          "Local Tennis Court",
// 		Description:       "Welcome to this awesome tennis match!",
// 		ParticipationFee:  0,
// 		RequiredEquipment: []string{"Racket", "Shoes"},
// 		Level:             "Any",
// 		ChatLink:          "https://example.com",
// 		StartsAt:          time.Date(2025, time.January, 1, 10, 0, 0, 0, time.UTC),
// 		EndsAt:            time.Date(2025, time.January, 1, 11, 30, 0, 0, time.UTC),
// 	}

// 	dbMatch := createData.Match()
// 	dbMatch.HostUserID = "DemoUser"

// 	return &dbMatch
// }

// func match2() *model.Match {
// 	minParticipants := int32(2)
// 	maxParticipants := int32(4)
// 	createData := model.MatchCreate{
// 		Sport:             "Badminton",
// 		MinParticipants:   &minParticipants,
// 		MaxParticipants:   &maxParticipants,
// 		Location:          "Sports Hall, Downtown",
// 		Description:       "Looking for people to play badminton with :)",
// 		ParticipationFee:  1000,
// 		RequiredEquipment: []string{},
// 		Level:             "Any",
// 		ChatLink:          "https://example.com",
// 		StartsAt:          time.Date(2025, time.January, 10, 18, 0, 0, 0, time.UTC),
// 		EndsAt:            time.Date(2025, time.January, 10, 19, 0, 0, 0, time.UTC),
// 	}

// 	dbMatch := createData.Match()
// 	dbMatch.HostUserID = "DemoUser"

// 	return &dbMatch
// }

// func createInitialData() {
// 	ctx := context.Background()

// 	m := dal.Q.Match
// 	p := dal.Q.Participation

// 	dbMatch1 := match1()
// 	dbMatch2 := match2()

// 	if err := m.WithContext(ctx).Create(dbMatch1, dbMatch2); err != nil {
// 		stdlog.Fatalf("Failed to create initial matches: %s", err.Error())
// 	}

// 	dbParticipation1 := model.Participation{MatchID: dbMatch1.ID, UserID: dbMatch1.HostUserID}
// 	if err := p.WithContext(ctx).Create(&dbParticipation1); err != nil {
// 		stdlog.Fatalf("Failed to create initial participation 1: %s", err.Error())
// 	}

// 	dbParticipation2 := model.Participation{MatchID: dbMatch2.ID, UserID: dbMatch2.HostUserID}
// 	if err := p.WithContext(ctx).Create(&dbParticipation2); err != nil {
// 		stdlog.Fatalf("Failed to create initial participation 2: %s", err.Error())
// 	}

// 	stdlog.Println("Data created")
// }
