package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	utils "github.com/bmorrisondev/go-utils"
	"github.com/golang-jwt/jwt"
	"github.com/tweetyah/lib"
	"github.com/tweetyah/lib/constants"
)

func main() {
	router := lib.NetlifyRouter{
		Get:    Get,
		Post:   Post,
		Put:    Put,
		Delete: Delete,
	}
	lambda.Start(router.Handler)
}

func Get(request events.APIGatewayProxyRequest, claims jwt.MapClaims, db *sql.DB) (events.APIGatewayProxyResponse, error) {
	var posts []lib.Post

	filter := request.QueryStringParameters["filter"]

	if filter == "scheduled" {
		query := "select id, text, send_at, status, linked_media from posts where id_user = ? and status = ?"
		res, err := db.Query(query, claims["user_id"], constants.PostStatus_Scheduled)
		if err != nil {
			return utils.ErrorResponse(err, "query db")
		}

		for res.Next() {
			var p lib.Post
			var mediaString string
			err = res.Scan(&p.Id, &p.Text, &p.SendAt, &p.Status, &mediaString)
			if err != nil {
				return utils.ErrorResponse(err, "scan")
			}
			if mediaString != "" {
				err := json.Unmarshal([]byte(mediaString), &p.Media)
				if err != nil {
					return utils.ErrorResponse(err, "unmarshal media")
				}
			}
			posts = append(posts, p)
		}
	} else if filter == "sent" {
		query := "select id, text, send_at, status from posts where id_user = ? and status = ?"
		res, err := db.Query(query, claims["user_id"], constants.PostStatus_Sent)
		if err != nil {
			return utils.ErrorResponse(err, "query db")
		}

		for res.Next() {
			var p lib.Post
			err = res.Scan(&p.Id, &p.Text, &p.SendAt, &p.Status)
			if err != nil {
				return utils.ErrorResponse(err, "scan")
			}
			posts = append(posts, p)
		}
	} else {
		query := "select id, text, send_at, status from posts where id_user = ?"
		res, err := db.Query(query, claims["user_id"])
		if err != nil {
			return utils.ErrorResponse(err, "query db")
		}

		for res.Next() {
			var p lib.Post
			err = res.Scan(&p.Id, &p.Text, &p.SendAt, &p.Status)
			if err != nil {
				return utils.ErrorResponse(err, "scan")
			}
			posts = append(posts, p)
		}
	}

	jstr, err := utils.ConvertToJsonString(posts)
	if err != nil {
		return utils.ErrorResponse(err, "conver to json string")
	}

	return utils.OkResponse(&jstr)
}

func Post(request events.APIGatewayProxyRequest, claims jwt.MapClaims, db *sql.DB) (events.APIGatewayProxyResponse, error) {
	var posts []lib.Post
	err := json.Unmarshal([]byte(request.Body), &posts)
	if err != nil {
		return utils.ErrorResponse(err, "json.Unmarshal")
	}

	log.Println(claims)

	userId := claims["user_id"].(string)
	serviceId := claims["service_id"].(string)
	serviceIdNum, err := strconv.Atoi(serviceId)
	if err != nil {
		return utils.ErrorResponse(err, "(Post) cast service id to num")
	}
	userIdNum, err := strconv.Atoi(userId)
	if err != nil {
		return utils.ErrorResponse(err, "(Post) cast user id to num")
	}

	if len(posts) == 1 {
		updated, err := lib.SavePostToDb(userIdNum, serviceIdNum, posts[0])
		if err != nil {
			return utils.ErrorResponse(err, "(Post) save post to db")
		}

		jstr, err := utils.ConvertToJsonString(updated)
		if err != nil {
			return utils.ErrorResponse(err, "utils.ConvertToJsonString")
		}
		return utils.OkResponse(&jstr)
	} else {
		threadStart, err := lib.SaveThreadToDb(userIdNum, serviceIdNum, posts)
		if err != nil {
			return utils.ErrorResponse(err, "(Post) save thread to db")
		}

		jstr, err := utils.ConvertToJsonString(threadStart)
		if err != nil {
			return utils.ErrorResponse(err, "(Post) utils.ConvertToJsonString")
		}
		return utils.OkResponse(&jstr)
	}
}

func Put(request events.APIGatewayProxyRequest, claims jwt.MapClaims, db *sql.DB) (events.APIGatewayProxyResponse, error) {
	var posts []lib.Post
	err := json.Unmarshal([]byte(request.Body), &posts)
	if err != nil {
		return utils.ErrorResponse(err, "json.Unmarshal")
	}

	userId := claims["user_id"].(string)
	serviceId := claims["service_id"].(string)
	serviceIdNum, err := strconv.Atoi(serviceId)
	if err != nil {
		return utils.ErrorResponse(err, "(Put) cast service id to num")
	}
	userIdNum, err := strconv.Atoi(userId)
	if err != nil {
		return utils.ErrorResponse(err, "(Put) cast user id to num")
	}

	if len(posts) == 1 {
		err = lib.UpdatePostInDb(userIdNum, serviceIdNum, posts[0])
		if err != nil {
			return utils.ErrorResponse(err, "(Put) save post to db")
		}

		return utils.OkResponse(nil)
	} else {
		err = lib.UpdateThreadInDb(userIdNum, serviceIdNum, posts)
		if err != nil {
			return utils.ErrorResponse(err, "(Put) save thread to db")
		}

		return utils.OkResponse(nil)
	}
}

func Delete(request events.APIGatewayProxyRequest, claims jwt.MapClaims, db *sql.DB) (events.APIGatewayProxyResponse, error) {
	var posts []lib.Post
	err := json.Unmarshal([]byte(request.Body), &posts)
	if err != nil {
		return utils.ErrorResponse(err, "json.Unmarshal")
	}

	userId := claims["user_id"].(string)
	serviceId := claims["service_id"].(string)
	serviceIdNum, err := strconv.Atoi(serviceId)
	if err != nil {
		return utils.ErrorResponse(err, "(Delete) cast service id to num")
	}
	userIdNum, err := strconv.Atoi(userId)
	if err != nil {
		return utils.ErrorResponse(err, "(Delete) cast user id to num")
	}

	err = lib.DeletePostsFromDb(userIdNum, serviceIdNum, posts)
	if err != nil {
		return utils.ErrorResponse(err, "(Delete) save thread to db")
	}

	return utils.OkResponse(nil)
}
