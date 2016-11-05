## Masnun Bot 
My personal Telegram bot 

#### Goal 
The key objective is to create some sort of automated notification system to stay up to date with various 
information sources (reddit, hacker news, rss feeds, tweets etc.) from one place. 

The bot will gather stuff from various sources and push them to me on Telegram. All sorts of contents mashed 
up in one place. 


#### Setup

The different credentials and configuration values would be stored as environment variables. I have a 
`configure.sh` file which exports the variables. The project includes a `configure.sh.sample` which you 
can use as a template. I am using Glide to manage the project dependencies. I recommend you use it too. 
You can install it from here: https://glide.sh/ 

The following steps should set the project up - 


* `mv configure.sh.sample configure.sh`
* Edit `configure.sh` to update your credentials / details
* `. ./configure.sh`
* `glide install`
* `go build`
* `./masnun-bot`