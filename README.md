# Movie-API

# Write your query or mutation here
mutation createMovie {
	createMovie(input: { title: "Shawshank Redemption", url: "https://www.imdb.com/title/tt0111161/" , directorId: "1"})
  {
  	title
  	url
	  author
    {
      id
    }
	}
}

# Write your query or mutation here
mutation createMovie {
	createMovie(input: { title: "Good Will Hunting", url: "https://www.imdb.com/title/tt0119217/" , directorId: "2"})
  {
  	title
  	url
	  author
    {
      id
    }
	}
}

query findMovies{
  movies{
    title
    url
    author{
      name
    }
  }
}

To run, use command 'go run server.go' from root directory
