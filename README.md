# RapidAPI

## RapidAPI is a collection of API's that RapidAPI has.  I wanted to build a hub, and practice consuming multiple types, as well as having multiple endpoints.  I will provide documention for each collection.
---


1. Covid19

```
/covid19
```

### The Covid19 collection takes in a country, date, and type; then returns the totals for the specified input.
|term|input|required|
|:--:|:---:|:------:|
|country||yes|
|date|mm-dd-yyyy|yes|
|type|confirmed, recovered, death|yes|

Example:

`/covid19?country=us&date=07-14-2020&countType=confirmed`

Result:

```
Retrieving confirmed count for us on 07-14-2020
On 07-14-2020; US declared 3431574 total CONFIRMED cases of Covid-19
```
