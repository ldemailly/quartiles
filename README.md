# quartiles
Find quartiles (apple news+ new game)


Make a screenshot, select the word fragments, paste to textedit, convert to plain text, paste to `quartiles`

Example:
```
$ go run . -loglevel debug "nder
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
10:40:06.232 [INF] Log level is now 0 Debug (was 2 Info)
10:40:06.274 [INF] Words in dictionary /usr/share/dict/words: 234456
10:40:06.274 [INF] Looking at fragments:
nder	nt	sti	bi
lo	thu	al	con
ies	vo	mon	ba
tue	rd	fri	lit
sw	ck	wed	gam
10:40:06.274 [INF] Checking 2 words combinations:
alba
albi
almon
baal
babi
back
bacon
bander
bant
bard
bick
bilo
binder
bint
bird
bisti
conal
gamba
gammon
lock
lord
monal
stick
stint
thunder
vowed
10:40:06.274 [INF] Checking 3 words combinations:
stibial
swallo
wedlock
10:40:06.276 [INF] Checking 4 words combinations:
backgammon
birdbander
constituent
thunderbird
thunderstick
```
