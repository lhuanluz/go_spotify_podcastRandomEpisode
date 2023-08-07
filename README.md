# Spotify Random Podcast Episode Opener

This Go program fetches all episodes from a specified Spotify podcast and opens a random episode in the Spotify desktop application. Configuration details like the Spotify client ID, client secret, and podcast ID are stored in a `.env` file.

## Prerequisites

1.  You need to have Go installed. You can [download it here](https://golang.org/dl/).
2.  Spotify desktop application installed on your machine.
3.  Developer credentials (Client ID and Client Secret) from the Spotify Developer Dashboard. If you haven't registered your app yet, you can do it [here](https://developer.spotify.com/dashboard/applications).

## Setup and Usage

### Step 1: Clone the Repository
`git clone https://github.com/lhuanluz/go_spotify_podcastRandomEpisode
cd go_spotify_podcastRandomEpisode` 

### Step 2: Install Dependencies

Install the required Go package:
`go get github.com/joho/godotenv` 

### Step 3: Set Up the .env File

In the root directory of the project, create a file named `.env`.

Add the following content:


`CLIENT_ID=YOUR_SPOTIFY_CLIENT_ID
CLIENT_SECRET=YOUR_SPOTIFY_CLIENT_SECRET
PODCAST_ID=YOUR_SPOTIFY_PODCAST_ID` 

Replace `YOUR_SPOTIFY_CLIENT_ID`, `YOUR_SPOTIFY_CLIENT_SECRET`, and `YOUR_SPOTIFY_PODCAST_ID` with your actual Spotify credentials and the ID of the podcast you're interested in.


### Step 4: Run the Program
Execute the program:
`go run main.go` 

If everything is set up correctly, the program will fetch a random episode from the specified podcast and open it in your Spotify desktop application.

## Troubleshooting

1.  Ensure your `.env` file is correctly set up and in the root directory of the project.
2.  Ensure the Spotify app is installed on your machine.
3.  If you encounter errors related to fetching episodes, print out the API responses to see potential issues.
4.  Ensure your Spotify developer credentials have the necessary permissions.

## Conclusion

With this Go program, you can easily surprise yourself with a random episode from your favorite podcast on Spotify. Remember to always keep your client secret and other sensitive details secure.
