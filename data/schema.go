package data

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
)

var rawPhotoType *graphql.Object
var userType *graphql.Object

var photoConnection *relay.GraphQLConnectionDefinitions

var nodeDefinitions *relay.NodeDefinitions

// Schema is our published GraphQL representation of objects and mutations
var Schema graphql.Schema

func init() {
	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ct context.Context) (interface{}, error) {
			resolvedID := relay.FromGlobalID(id)
			if resolvedID.Type == "RawPhoto" {
				return GetRawPhoto(resolvedID.ID), nil
			} else if resolvedID.Type == "User" {
				return GetUser(resolvedID.ID), nil
			}
			return nil, nil
		},
		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			switch p.Value.(type) {
			case *RawPhoto:
				return rawPhotoType
			case *User:
				return userType
			}
			return nil
		},
	})

	rawPhotoType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "RawPhoto",
		Description: "A photo binary blob",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("RawPhoto", nil),
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})

	photoConnection = relay.ConnectionDefinitions(relay.ConnectionConfig{
		Name:     "PhotoConnection",
		NodeType: rawPhotoType,
	})

	userType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "A user",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("User", nil),
			"photos": &graphql.Field{
				Type:        photoConnection.ConnectionType,
				Description: "A user's collection of photos",
				Args:        relay.ConnectionArgs,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := relay.NewConnectionArguments(p.Args)
					dataSlice := PhotosToInterfaceSlice(GetRawPhotos()...)
					return relay.ConnectionFromArray(dataSlice, args), nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"node": nodeDefinitions.NodeField,

			// Add you own root fields here
			"viewer": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetViewer(), nil
				},
			},
		},
	})

	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		// Mutation: mutationType,
		Query: queryType,
		Types: []graphql.Type{queryType, userType},
	})
	if err != nil {
		panic(err)
	}
}

// PhotosToInterfaceSlice gets an interface slice.
// See https://github.com/golang/go/wiki/InterfaceSlice
func PhotosToInterfaceSlice(photos ...*RawPhoto) []interface{} {
	var interfaceSlice = make([]interface{}, len(photos))
	for i, d := range photos {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}
