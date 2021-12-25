# Google text-to-speak Simple Binary file

This repository is a simple implementation of google text-to-speak service. \
Required enable API in GCP (https://console.cloud.google.com/apis/api/texttospeech.googleapis.com)


## Usage 

First and only argument is text for speak ;) \
You can change behavior of binary by flags:

|Option           |Default     |Description|
|------------------|------------|-----------|
|credentials-path |cloud-proxy |_Path to credential file from GCP, default GCP cloud proxy auth_|
|output-path      |output.mp3  |_Path to output file with mp3 footage_|
|pl              |false       |_Flag, when would you to use polish language_|

### 1. The simplest usage

Generate file `output.mp3` in workdir in English with GCP Cloud Proxy Auth:

```bash
google-speak "Welcome! How are you?"
```

### 2. With polish and custom auth file

Generate file `output.mp3` in workdir in Polish with `credentials.json` auth file.

```bash
google-speak -pl -credentials-path credentials.json "Witaj! Co słychać?"
```

### Help message (`-h`)

```bash
Usage of ./google-speak:
  -credentials-path string
        Path to credential file (JSON) (default "cloud-proxy")
  -output-path string
        Output path (mp3 ext required) (default "output.mp3")
  -pl
        Change context to polish language
```



## Installation

Simple pre-compiled version:

1. Download binary from releases (https://github.com/oxess/google-speak/releases)
2. Upload to server filesystem `scp google-speak server1:/usr/local/bin/google-speak`  
3. Make executable `chmod +x /usr/local/bin/google-speak`

Or:

1. Download repository `git clone https://github.com/oxess/google-speak`
2. Go to directory `cd google-speak`
3. Build dependencies `go mod download`
4. Compile to binary `go build`
5. Move to local filesystem `mv google-speak /usr/local/bin/google-speak`
6. Can file execute `chmod +x /usr/local/bin/google-speak`



## How to generate GCP JSON Access Key

1. Enable text-to-speech API in GCP (https://console.cloud.google.com/apis/api/texttospeech.googleapis.com)
2. Create new service account (https://console.cloud.google.com/iam-admin/serviceaccounts/create)
3. Go to detail view of this service account > Keys
4. Push button "Add key" and "Create new key" and choose "JSON"
5. Save file on disk
6. Done!

> **Attention!** \
> You should secure your key for production version. \
> Do not using didn't secure key on production!
