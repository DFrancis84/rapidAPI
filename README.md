# RapidAPI

## RapidAPI is a collection of API's that RapidAPI has.  I wanted to build a hub, and practice consuming multiple types, as well as having multiple endpoints.  I will provide documention for each collection.
---

_You will need to visit https://rapidapi.com/ to get your api key that you will need to store in your environment variables as "X_RAPIDAPI_KEY"_

---


**1. Covid19**
## This collection accesses data from John Hopkins based on the type of cases you are researching.  (Confirmmed, Recovered or Death).

```
/covid19
```

### The Covid19 collection takes in a country, date, and type; then returns the totals for the specified input.
|field|input|required|
|:--:|:---:|:------:|
|country|[list of valid countries](internal/covid19/validCountries.txt)|yes|
|date|mm-dd-yyyy|yes|
|type|confirmed, recovered, death|yes|

Example:

`/covid19?country=us&date=07-14-2020&countType=confirmed`

Result:

```
Retrieving confirmed count for us on 07-14-2020
On 07-14-2020; US declared 3431574 total CONFIRMED cases of Covid-19
```

---

**2. Urban Dictionary**
## This collection accesses the Urban Dictionary, and returns their definition for the term passed through.

```
/urbandictionary
```

### The UrbanDictionary collection takes in a term; then returns the dictionary specified by UD
|field|input|required|
|:--:|:---:|:------:|
|term|any word|yes|

Example:

`/urbandictionary?term=slap`

Result:

**WARNING: The results returned from UD CAN contain offensive or vulgar language.  Please use at your own risk**

```
Searching Urban Dictionary for 'slap'
1:      [What the] [five fingers] said [to the face].
2:      a group of *BASS* lovers, a fellow followers of [Davie504], they [slap the bass], [like button], even coming to your house to slap your *BASS*
3:      Slap is tight music, somthing you can go dumb to.

Slap is what a pimp says to a hoes face when she shows up funny wit the ends.

Slap is Having 10s 12s or 15s in your trunk [nocking] while your driving,
making you trunk [raddle].

This word is from th BAY AREA, California. Where we do the Dam' thing. [YaDidiMean]....

and for all you fake white boys quit trying to be black, bitch ass marks....
4:      [Sounds like a plan]!
5:      To have a very loud, good quality sound system esp. in your car. Or a song that just sounds hella good and has tremendous bass in it. I have to set the record [str8] once again. We started this in the Town, Oakland, Ca, USA. When everybody else was still callin it "Knock" or "[Thump]" We in Oakland started calling it "Slap". The rest of the Bay Area eventually caught on to it. [Str8] Up!
6:      [stands] for [Sounds Like A Plan].
7:      [sounds like a plan]
8:      [A slap] is a [coded] way of saying vape. When someone says can I get [a slap] they are asking for a hit. This term comes from the sound that a guy could make using his hand, it sounded like a slap and that's how he asked for hits of a vape so he would not be caught by staff/teachers.
9:      Something you give someone when you aren't quite [pissed off] enough to give them a [Krotch] Kick/[Cunt Punt].
10:     Verb.

1. Used to describe giving or taking something.
2. To [slap it] [somewhere] - Meaning to go somewhere
3. Slap something up - To [create]

```
_The formatting on the output from UD is a little off as I can separate each individual definition, but each definition may have a couple "sub definitions" that throw it off_
