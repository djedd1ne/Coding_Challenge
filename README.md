# Coding Challenge
### Get TELERAM_BOT_TOKEN and TELEGRAM_CHAT_ID 
1. Find telegram bot named "@botfarther"
3. To create a new bot type `/newbot`
4. Choose a name for your bot, then a username that must end with "bot"
5. Copy your API token to your .env file
6. type `/setprivacy`, then select your bot
7. Choose `Disable` to allow the bot to read all messages in the group
Then, create a group on telegram,
8. Add your bot to the group
9. Send a test message in the group
10. Run
    ```shell
    curl -X GET "https://api.telegram.org/bot<YOUR_BOT_TOKEN>/getUpdates"
    ```
    you will see
    ```json
    "chat":{"id":-123456789,"title":"Group Name"}
    ```
    in the response
12. Copy the negative value to your .env file
13. Your .env file must look like this 
```.env
TELEGRAM_BOT_TOKEN=YourApiToken
TELEGRAM_CHAT_ID=YourGroupId
```

### Clone and run
Clone the repository
```shell
git clone https://github.com/djedd1ne/Coding_Challenge.git
```
Navigate into the directory
```shell
cd Coding_Challenge/
```
Build
```shell
go build -o bin/server ./cmd/server
```
Run
```shell
./bin/server
```
### Test
1. Using curl (application listening on port 5000)
* Send a valid notificaton
```shell
curl -X POST http://localhost:5000/notify \
     -H "Content-Type: application/json" \
     -d '{"Type": "Warning", "Name": "Backup Failure", "Description": "The backup failed due to a database problem"}'
```
2. Using Postman
