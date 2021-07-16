# discord-bsc-gas-price

⚠️This project has been merged with [discord stock tickers](https://github.com/rssnyder/discord-stock-ticker).⚠️

a discord bot to display the current reccomended gas amount for binance smart chain transactions

![Live Demo Gif](https://s3.cloud.rileysnyder.org/public/assets/bscgas.gif)

## Support
<a href='https://ko-fi.com/rileysnyder' target='_blank'><img height='35' style='border:0px;height:46px;' src='https://az743702.vo.msecnd.net/cdn/kofi3.png?v=0' border='0' alt='Buy Me a Coffee' />
[![Discord Chat](https://img.shields.io/discord/806606291798982678)](https://discord.gg/CQqnCYEtG7)

## Add to your discord server (click image)

[![Gas Price](https://s3.cloud.rileysnyder.org/public/assets/bscgas.png)](https://discord.com/api/oauth2/authorize?client_id=856947934452645898&permissions=0&scope=bot)

## Self Host

Self hosting allows you to update the price in the bot name rather than the activity on the public versions.

### Binary

Download the latest binary from the [release page](https://github.com/rssnyder/discord-bsc-gas-price/releases).

Or build from source:

```
git clone https://github.com/rssnyder/discord-bsc-gas-price.git

cd discord-bsc-gas-price

go build -o discord-bsc-gas-price
```

Create a bot in the discord dev portal, and grab the token. When you add it to your server, be sure to give it "Change Nickname" permissions.

Run the bot with the nessesary settings:

```
  -frequency int
        seconds between gas price cycles (default 5)
  -setNickname
        wether to set nickname of bot
  -token string
        discord bot token
```

```
./discord-bsc-gas-price -token 'xxxxxxxxxxx' -setNickname
```

You can also use the template `discord-bsc-gas-price.service` file to install it as a systemd service.

Fill in your discord bot token in the file, and add any extra arguments you need.

```
cp discord-bsc-gas-price.service /etc/systemd/system/

mkdir -p /etc/discord-bsc-gas-price

cp discord-bsc-gas-price /etc/discord-bsc-gas-price/

systemctl daemon-reload 

systemctl start discord-bsc-gas-price.service 
```
