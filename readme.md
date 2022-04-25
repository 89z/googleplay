# GooglePlay

> Fear plays an interesting role in our lives. How dare we let it motivate us?
> How dare we let it into our decision-making, into our livelihoods, into our
> relationships?
>
> It's funny, isn't it, we take a day a year to dress up in costume and
> celebrate fear?
>
> [Spooked (The Office) (2011)][1]

Download APK from Google Play or send API requests

[1]://f002.backblazeb2.com/file/ql8mlh/Spooked+%28The+Office%29.mp4

## How to install?

This module works with Windows, macOS or Linux. You can download, build and run
in less than [one&nbsp;minute][2].

First, [download Go][3] and extract archive. Note that you need go >= 1.18. Then [download GooglePlay][4] and
extract archive. Then navigate to `googleplay-master/cmd/googleplay`, and
enter:

~~~
go build
~~~

[2]://f002.backblazeb2.com/file/ql8mlh/googleplay.webm
[3]://go.dev/dl
[4]://github.com/89z/googleplay/archive/refs/heads/master.zip

## Tool examples

If you are outside the United States, you might need to create an
[App Password][5]. Create a file containing token (`aas_et`) for future
requests:

~~~
googleplay -e EMAIL -p PASSWORD
~~~

Create a file containing `Android_ID` (GSF ID) for future requests:

~~~
googleplay -d
~~~

Get app details:

~~~
PS C:\> googleplay -a com.google.android.youtube
Title: YouTube
Creator: Google LLC
UploadDate: Mar 16, 2022
VersionString: 17.11.34
VersionCode: 1528288704
NumDownloads: 11.562 B
Size: 40.935 MB
Files: 4
Offer: 0 USD
~~~

Purchase app. Only needs to be done once per Google account:

~~~
googleplay -a com.google.android.youtube -purchase
~~~

Download APK. You need to specify any valid version code. The latest code is
provided by the previous details command. If APK is split, all pieces will be
downloaded:

~~~
googleplay -a com.google.android.youtube -v 1528288704
~~~

[5]://support.google.com/accounts/answer/185833

## Sponsor

I really need help financially, so if you are able, please donate using the
sponsor link. If you cannot use Paypal, let me know, and I can see about adding
other methods. Thank you!

## API

https://github.com/89z/googleplay
