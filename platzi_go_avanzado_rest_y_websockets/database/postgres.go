package database

import (
	"context"
	"database/sql"
	"log"
	"platzi/go/rest_websockets/models"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (p *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := p.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password)
	return err
}

func (p *PostgresRepository) InsertPost(ctx context.Context, post *models.Post) error {
	_, err := p.db.ExecContext(ctx, "INSERT INTO posts (id, user_id, post_content) VALUES ($1, $2, $3)", post.Id, post.UserId, post.PostContent)
	return err
}

func (p *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	rows, _ := p.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	rows, _ := p.db.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", email)

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Email, &user.Password); err == nil {
			return &user, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *PostgresRepository) GetPostById(ctx context.Context, id string) (*models.Post, error) {
	rows, _ := p.db.QueryContext(ctx, "SELECT id, user_id, post_content, created_at FROM posts WHERE id = $1", id)

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var post models.Post
	for rows.Next() {
		if err := rows.Scan(&post.Id, &post.UserId, &post.PostContent, &post.CreatedAt); err == nil {
			return &post, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &post, nil
}

func (p *PostgresRepository) UpdatePost(ctx context.Context, post *models.Post) error {
	_, err := p.db.ExecContext(ctx, "UPDATE posts SET post_content = $1 WHERE id = $2 AND user_id = $3", post.PostContent, post.Id, post.UserId)
	return err
}

func (p *PostgresRepository) DeletePost(ctx context.Context, id string, userId string) error {
	_, err := p.db.ExecContext(ctx, "DELETE FROM posts WHERE id = $1 AND user_id = $2", id, userId)
	return err
}

func (p *PostgresRepository) ListPosts(ctx context.Context, page uint64) ([]*models.Post, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, user_id, post_content, created_at FROM posts LIMIT $1 OFFSET $2", 2, page*2)

	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Id, &post.UserId, &post.PostContent, &post.CreatedAt); err == nil {
			posts = append(posts, &post)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *PostgresRepository) Close() {
	p.db.Close()
}

//// docker build . -t platzi-ws-rest-db
//// docker run -p 54321:5432 platzi-ws-rest-db  //este es el comando para inicar base de datos
//// docker build . -t platzi-ws-rest-db
