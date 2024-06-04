# Quartiles

Find quartiles ([apple news+ new game](https://www.apple.com/newsroom/2024/05/apple-news-plus-introduces-quartiles-a-new-game-and-offline-mode-for-subscribers/))

You can make a screenshot and use iphone text extraction to copy paste the list of fragments (or type it)

## Installation

From source
```
go install github.com/ldemailly/quartiles@latest
```

Or see the numerous binaries in https://github.com/ldemailly/quartiles/releases

Or docker `fortio/quartiles:latest`

Or brew `brew install fortio/tap/quartiles`

(I manage the fortio org and usually put everything there but this one is completely unrelated outside of it's usage of https://pkg.go.dev/fortio.org/sets#Tuplets so for now it is hosted here in `ldemailly` yet uses fortio's org for docker and brew)

## Example

Make a screenshot, select the word fragments, paste to textedit, convert to plain text, paste to `quartiles`

Example:
```
$ docker run -ti fortio/quartiles:latest "nder
nt
sti
bi
lo
thu
al
con
ies
VO
mon
ba
tue
rd
fri
lit
SW
ck
wed
gam
"
```

```
19:22:47.346 [INF] Words in dictionary /usr/share/dict/words: 40358
19:22:47.346 [INF] Looking at fragments:
nder	nt	sti	bi
lo	thu	al	con
ies	vo	mon	ba
tue	rd	fri	lit
sw	ck	wed	gam
19:22:47.346 [INF] Checking 2 words combinations:
back
bacon
bard
binder
bird
lock
lord
lowed
monies
stick
stint
thunder
vowed
19:22:47.346 [INF] Checking 3 words combinations:
allowed
stickies
wedlock
19:22:47.351 [INF] Checking 4 words combinations:
backgammon
constituent
frivolities
swallowed
```

## Usage

```
quartiles 0.1.0 usage:
	quartiles [flags] word fragments...
Finds all words in dictionary that can be made from the given fragments
or 1 of the special arguments
	quartiles {help|envhelp|version|buildinfo}
flags:
  -dictionary path
    	Dictionary file path to use (default "/usr/share/dict/words")
```

## What is the Quartiles game?

https://www.apple.com/newsroom/2024/05/apple-news-plus-introduces-quartiles-a-new-game-and-offline-mode-for-subscribers/


<img src="https://github.com/ldemailly/quartiles/assets/3664595/4ef37a1b-86e8-4841-88fa-c36bc5e5838b" width="50%">

## Wordlist

The `words` used in Dockerfile is from Ubuntu's
`apt install wamerican-large` american-english-large words-large
removing 's and all caps acronyms
```
grep -v "'" /usr/share/dict/american-english-large | grep -v -E "^[A-Z]+s?$" > words
```
