# Local Music Discovery Hub

A passion project created and designed for local artists, bands, labels, and venues 
to pull together a whole music community into one central hub.

I have designed this as a music producer and DJ myself, I found there wasn't really 
a standardised way to find and explore the local music scene for myself and others, as the only 
methods I and many other people currently use are through ticketing platforms, labels, and social media etc. 
## Stack

### *Front-end*
* **Language**: JavaScript/TypeScript
* **Frameworks**: Astro, React, Tailwind CSS

### *Backend*
* **Language**: Go ('net/http`)
* **Database**: SQLite ('modernc.org/sqlite`)

## Endpoints

### Venues
* **GET `/venues`**
  * Returns a JSON list of registered local venues.
  * *Query Parameters*: `region` (string), `location` (string)

### Labels
* **GET `/labels`**
  * Returns a JSON list of registered local labels.
  * *Query Parameters*: `name` (string), `location` (string), `region` (string), `genre` (string), `link`, (string)

### Artists
* **GET `/artists`**
  * Returns a JSON list of local artists and DJ's
  * *Query Parameters*: `name` (string), `dj_name` (string), `location` (string), `region` (string), `genre` (string), `link`, (string)

### Genres
* **GET `/genres`**
  * Returns a JSON list existing genres.
  * *Query Parameters*: `name` (string)

## Development & Learning



---

## Disclaimer & Community Notice

* **Non-Commercial Project**: This platform is a non-profit, open-source community directory built purely to spotlight Wollongong, The Illawarra and Sydney talent (Possibly more in future).
* **Data & Attribution**: All artist, venue, and label details are gathered from public platforms and social channels.
* **Updates & Removals**: If you are an artist, venue owner, or label manager featured here and would like your information updated or removed, please open an issue or submit a pull request.
