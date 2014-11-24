tsumtsum
========

Contact manager for LINE messaging service, specifically useful for LINE: Disney Tsum Tsum for inviting. 

###Quick start
Install: 
```
go get github.com/bcgraham/tsumtsum
```
Or download: 
    <a href="http://itwill.be/compiled/tsumtsum">Mac OS X</a> (SHA: 21c041196d27544c155dfeb1e7e1997f116a202b) 
    <a href="http://itwill.be/compiled/tsumtsum.exe">Windows</a> (SHA: 02c6b888687091ff0e59f1798845db28ef2068ab)

First-time user? Want to get going now? This is probably what you want to do. Replace LINELOGIN with your LINE login name. 
```
$ tsumtsum -u=LINELOGIN whitelist -o=whitelist.json
$ tsumtsum -u=LINELOGIN add 
```
###Features
This is a command-line tool for adding adding and removing contacts from your LINE contact list. Its main features are:
* It works on **Windows, OS X, and Linux**. 
* You can use your computer while it **runs in the background**, because it's run from the terminal. 
* It's a standalone utility. It doesn't care if the LINE app is installed on your computer. 
* It's **fast**. Adding 500 contacts takes about 3 minutes. 
* It can purge your contacts list even faster. 
* Keep your actual LINE contacts by **making a whitelist**. 
* By default, it pulls new user IDs from a service I have running. My service has about **15,000 user IDs** in it, a lot of which are new - enough to keep you busy for a while.
* When adding IDs, they report to the service whether the ID was found or not. After an ID isn't found a few times, it's retired from distribution. This means my lists run a hit rate of about 90% - not like stale lists of user IDs floating around in text files. 
* It can add more users than the desktop LINE clients. It can be finicky - I've been blocked after 500, but I've also added thousands in one day. I cleared **2 million coins in one day** using this tool.

###Usage

There are a few options that are relevant to all subcommands. These are:

* -u : Your LINE login name. 
* -d : Your device name. This defaults to "PC" because I have never seen a value rejected, but it's an option if you need it to be. 
* -r : Reporting service. This defaults to the service I'm running. I've included the code for the service I'm running in the /server directory, in case you want to roll your own. 

Before you start, you should make a whitelist. 
```
$ tsumtsum -u=LINELOGIN whitelist -o=whitelist.json
```
The whitelist command takes an output file, in which it saves your current contact list. This will be useful, later, for purging your contacts list. 

After you've made your whitelist, it's time to get adding. 
```
$ tsumtsum -u=LINELOGIN add -s=userids.txt -l=750
```
Adding takes the following options: 

* -s : Source of user IDs. This defaults to my service, but it will also accept text files with user IDs separated by newlines (i.e., the format of all the lists of user IDs floating around). 
* -l : Limit. It defaults to 500. I've gotten temp-bans from as low as adding 600. The tempbans last between 12 and 24 hours - same as for the desktop LINE client. 

After adding, you'll want to invite everyone. While you don't have to, I like to leave the service running on "monitor". This listens to the invites you send and deletes the contacts after you've invited them. Purging is really, really fast (deletes 3-4 contacts per second) but it's also nice to keep a running tally of the invites you've sent so far. 
```
$ tsumtsum -u=LINELOGIN monitor 
```

After you're done inviting, you'll want to purge your contacts list. 
```
$ tsumtsum -u=LINELOGIN purge -w=whitelist.json
```
This deletes all your contacts except those you've saved on the whitelist. It's really fast. 

If you have any questions, feel free to shoot me an email at bcgraham@gmail.com . 
