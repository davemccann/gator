package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/davemccann/gator/internal/database"
)

func command_browse(s *state, cmd command, user *database.User) error {
	limit := 2
	if len(cmd.arguments) > 0 {
		val, err := strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return err
		}
		limit = val
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.dbQueries.GetPostsForUser(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("\n==== Displaying followed posts for user: %s ====\n", user.Name)
	for _, post := range posts {
		printPost(&post)
	}

	return nil
}

func printPost(post *database.GetPostsForUserRow) {
	outputFormat := `
%s | published: %s
    
Title: %s

Description: %s

Link: %s

===============================================================
`
	fmt.Printf(outputFormat, post.FeedName, post.PublishedAt, post.Title, post.Description, post.Url)
}
