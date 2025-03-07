---
title: Ranking View Query Results
menuTitle: Ranking
weight: 70
description: >-
  You can query Views and return the most relevant results first based on their ranking score
archetype: default
---
{{< description >}}

ArangoSearch supports the two most popular ranking schemes:

- [Okapi BM25](https://en.wikipedia.org/wiki/Okapi_BM25)
- [TF-IDF](https://en.wikipedia.org/wiki/Tf%E2%80%93idf)

Under the hood, both models rely on two main components:

- **Term frequency** (TF):
  in the simplest case defined as the number of times a term occurs in a document
- **Inverse document frequency** (IDF):
  a measure of how relevant a term is, i.e. whether the word is common or rare
  across all documents

See _Ranking in ArangoSearch_ in the
[ArangoSearch Tutorial](https://www.arangodb.com/learn/search/tutorial/#:~:text=Ranking%20in%20ArangoSearch)
to learn more about the ranking model.

## Basic Ranking

To sort View results from most relevant to least relevant, use a
[SORT operation](../../aql/high-level-operations/sort.md) with a call to a
[Scoring function](../../aql/functions/arangosearch.md#scoring-functions) as
expression and set the order to descending. Scoring functions expect the
document emitted by a `FOR … IN` loop that iterates over a View as first
argument.

```aql
FOR doc IN viewName
  SEARCH …
  SORT BM25(doc) DESC
  RETURN doc
```

You can also return the ranking score as part of the result.

```aql
FOR doc IN viewName
  SEARCH …
  RETURN MERGE(doc, { bm25: BM25(doc), tfidf: TFIDF(doc) })
```

Scoring functions cannot be used outside of `SEARCH` operations, as the scores
can only be computed in the context of a View, especially because of the
inverse document frequency (IDF).

### Dataset

[IMDB movie dataset](example-datasets.md#imdb-movie-dataset)

### View definition

#### `search-alias` View

```js
db.imdb_vertices.ensureIndex({
  name: "inv-text",
  type: "inverted",
  fields: [
    { name: "description", analyzer: "text_en" }
  ]
});

db._createView("imdb_alias", "search-alias", { indexes: [
  { collection: "imdb_vertices", index: "inv-text" }
] });
```

#### `arangosearch` View

```json
{
  "links": {
    "imdb_vertices": {
      "fields": {
        "description": {
          "analyzers": [
            "text_en"
          ]
        }
      }
    }
  }
}
```

### AQL queries

Search for movies with certain keywords in their description and rank the
results using the [`BM25()` function](../../aql/functions/arangosearch.md#bm25):

_`search-alias` View:_

```aql
FOR doc IN imdb_alias
  SEARCH doc.description IN TOKENS("amazing action world alien sci-fi science documental galaxy", "text_en")
  SORT BM25(doc) DESC
  LIMIT 10
  RETURN {
    title: doc.title,
    description: doc.description,
    score: BM25(doc)
  }
```

_`arangosearch` View:_

```aql
FOR doc IN imdb
  SEARCH ANALYZER(doc.description IN TOKENS("amazing action world alien sci-fi science documental galaxy", "text_en"), "text_en")
  SORT BM25(doc) DESC
  LIMIT 10
  RETURN {
    title: doc.title,
    description: doc.description,
    score: BM25(doc)
  }
```

| title | description | score |
|:------|:------------|:------|
| AVPR: Aliens vs. Predator - Requiem | Prepare for more mayhem as warring **aliens** and predators return … spectacular **action** sequences … | 35.85710525512695 |
| Moon 44 | … battle a familiar foe and an **alien** enemy. … **sci-fi** thriller from **action** director Roland Emmerich … | 35.85523223876953 |
| Dark Star | A low-budget, **sci-fi** satire … battle their **alien** mascot … | 28.655567169189453 |
| Starship Troopers 2: Hero of the Federation | In the sequel to Paul Verhoeven's loved/reviled **sci-fi** film … fighting **alien** bugs… | 28.635963439941406 |
| Push | The **action** packed **sci-fi** thriller involves a group of young American ex-pats… | 28.131816864013672 |
| Casshern | Live-action sci-fi movie based on a 1973 Japanese animé of the same name. | 28.070863723754883 |
| Puzzlehead | In a post apocalyptic **world** where technology is outlawed, … The resulting **Sci-Fi** love triangle is a Frankensteinian fable … | 25.57171630859375 |
| Cesta do pravěku | Most classical **sci-fi** from K. Zeman. … a wondrous prehistoric **world** … | 25.57117462158203 |
| Interstella 5555: The 5tory of the 5ecret 5tar 5ystem | A **sci-fi** japanimation House-musical movie … themes of **sci-fi** celebrity … | 22.481136322021484 |
| Alien Planet | The dynamic meeting of solid **science** … **Alien** Planet creates a realistic depiction of creatures on another **world**, … | 21.493724822998047 |

Do the same but with the [`TFIDF()` function](../../aql/functions/arangosearch.md#tfidf):

_`search-alias` View:_

```aql
FOR doc IN imdb_alias
  SEARCH doc.description IN TOKENS("amazing action world alien sci-fi science documental galaxy", "text_en")
  SORT TFIDF(doc) DESC
  LIMIT 10
  RETURN {
    title: doc.title,
    description: doc.description,
    score: TFIDF(doc)
  }
```

_`arangosearch` View:_

```aql
FOR doc IN imdb
  SEARCH ANALYZER(doc.description IN TOKENS("amazing action world alien sci-fi science documental galaxy", "text_en"), "text_en")
  SORT TFIDF(doc) DESC
  LIMIT 10
  RETURN {
    title: doc.title,
    description: doc.description,
    score: TFIDF(doc)
  }
```

| title | description | score |
|:------|:------------|:------|
| AVPR: Aliens vs. Predator - Requiem | Prepare for more mayhem as warring **aliens** and predators return … spectacular **action** sequences … | 25.193025588989258 |
| Moon 44 | … battle a familiar foe and an **alien** enemy. … **sci-fi** thriller from **action** director Roland Emmerich … | 25.193025588989258 |
| Interstella 5555: The 5tory of the 5ecret 5tar 5ystem | A **sci-fi** japanimation House-musical movie … themes of **sci-fi** celebrity … | 20.324928283691406 |
| Dark Star | A low-budget, **sci-fi** satire … battle their **alien** mascot … | 19.935544967651367 |
| Starship Troopers 2: Hero of the Federation | In the sequel to Paul Verhoeven's loved/reviled **sci-fi** film … fighting **alien** bugs… | 19.935544967651367 |
| Casshern | Live-action sci-fi movie based on a 1973 Japanese animé of the same name. | 19.629377365112305 |
| Push | The **action** packed **sci-fi** thriller involves a group of young American ex-pats… | 19.629377365112305 |
| Puzzlehead | In a post apocalyptic **world** where technology is outlawed, … The resulting **Sci-Fi** love triangle is a Frankensteinian fable … | 18.10955047607422 |
| Cesta do pravěku | Most classical **sci-fi** from K. Zeman. … a wondrous prehistoric **world** … | 18.10955047607422 |
| The Day the Earth Stood Still | An **alien** and a robot land on earth after **World** War II … A classic **science** fiction film … | 15.719740867614746

## Stable pagination for results

The `SORT` operation does not guarantee a stable sort if there is no unique value
to sort by. This leads to an undefined order when sorting equal documents.

If you run a query multiple times with varying `LIMIT` offsets for pagination,
you can miss results or get duplicate results if the sort order is undefined.
To achieve stable pagination, you must meet the following requirements:

- The dataset should not change between query runs.
- The `SORT` operation must have at least one field with a unique value to sort by.

When stable sort is required, you can use a tiebreaker field. If the application
has a preferred field that indicates the order of documents with the same score,
then this field should be used in the `SORT` operation as a tiebreaker. If there
is no such field, you can use the `_id` system attribute as it is unique and
present in every document. 

```aql
// sort by score and break ties using the unique document identifiers
SORT BM25(doc) DESC, doc._id
```

**Example**

You can use the [IMDB movie dataset](example-datasets.md#imdb-movie-dataset)
and create a [View](full-text-token-search.md#view-definition) for
the `imdb_vertices` collection and call it `imdb`. Index the `description`
attribute with the built-in `text-en` Analyzer. Then, you can run the following
query to find movies that contain the token `ninja` in the description, sorted
by best matching according to the Okapi BM25 scoring scheme:

```aql
FOR doc IN imdb
  SEARCH ANALYZER(doc.description IN TOKENS("ninja", "text_en"), "text_en")
  LET score = BM25(doc)
  SORT score DESC
  RETURN { title: doc.title, score }
```

Note the 5th and 6th result, which both have the same score of `6.30634880065918`:

| title | score |
|:------|:------|
| Red Shadow: Akakage | 8.882122039794922 |
| Beverly Hills Ninja | 7.128915786743164 |
| Naruto the Movie: Ninja Clash in the Land of Snow | 7.041049957275391 |
| TMNT | 7.002012729644775 |
| Teenage Mutant Ninja Turtles II: The Secret of the Ooze | **6.30634880065918** |
| Batman Begins | **6.30634880065918** |
| ... | ... |

If you add a `LIMIT` operation for pagination and fetch the first 5 results,
you may get the **Batman movie** as the 5th result:

```aql
FOR doc IN imdb
  SEARCH ANALYZER(doc.description IN TOKENS("ninja", "text_en"), "text_en")
  LET score = BM25(doc)
  SORT score DESC
  LIMIT 0, 5
  RETURN { title: doc.title, score }
```

| title | score |
|:------|:------|
| Red Shadow: Akakage | 8.882122039794922 |
| Beverly Hills Ninja | 7.128915786743164 | 
| Naruto the Movie: Ninja Clash in the Land of Snow | 7.041049957275391 | 
| TMNT | 7.002012729644775 | 
| **Batman Begins** | 6.30634880065918 | 

If you change the query to `LIMIT 5, 5` to get the second batch of results, then
you may see the **Batman movie again** (as the 6th result overall) instead of
the **Ninja Turtles movie**:

| title | score |
|:------|:------|
| **Batman Begins** | 6.30634880065918 | 
| Shogun Assassin  | 5.8539886474609375 | 
| Scooby-Doo and the Samurai Sword | 5.736422538757324 | 
| Revenge of the Ninja | 5.212964057922363 | 
| Winners & Sinners 2: My Lucky Stars | 5.165824890136719 | 

The problem is the undefined order of results with the same score. There is no
guarantee whether the Batman or the Ninja Turtle movie comes first, and as a
result, both batches may include the same movie and miss the other entirely.
As the order is undefined, you can randomly encounter this problem, even if it
seems to work at first, because it may work some of the time.

To establish a stable order, you change the query to `SORT score DESC, doc._id`.
This still ranks movies by the score, but matches with the same score are
consistently sorted by the document identifier to break ties. This guarantees
that either the Batman or the Ninja Turtles movie is included in the first batch
and the other movie in the second batch.

```aql
FOR doc IN imdb
  SEARCH ANALYZER(doc.description IN TOKENS("ninja", "text_en"), "text_en")
  LET score = BM25(doc)
  SORT score DESC, doc._id
  LIMIT 0, 5  // first batch
  RETURN { title: doc.title, score }
```

| title | score |
|:------|:------|
| Red Shadow: Akakage | 8.882122039794922 |
| Beverly Hills Ninja | 7.128915786743164 | 
| Naruto the Movie: Ninja Clash in the Land of Snow | 7.041049957275391 | 
| TMNT | 7.002012729644775 | 
| **Teenage Mutant Ninja Turtles II: The Secret of the Ooze** | 6.30634880065918 | 

```aql
FOR doc IN imdb
  SEARCH ANALYZER(doc.description IN TOKENS("ninja", "text_en"), "text_en")
  LET score = BM25(doc)
  SORT score DESC, doc._id
  LIMIT 5, 5  // second batch
  RETURN { title: doc.title, score }
```

| title | score |
|:------|:------|
| **Batman Begins** | 6.30634880065918 | 
| Shogun Assassin  | 5.8539886474609375 | 
| Scooby-Doo and the Samurai Sword | 5.736422538757324 | 
| Revenge of the Ninja | 5.212964057922363 | 
| Winners & Sinners 2: My Lucky Stars | 5.165824890136719 | 

## Query Time Relevance Tuning

You can fine-tune the scores computed by the Okapi BM25 and TF-IDF relevance
models at query time via the `BOOST()` AQL function and also calculate a custom
score. In addition, the `BM25()` function lets you adjust the coefficients at
query time.

The `BOOST()` function is similar to the `ANALYZER()` function in that it
accepts any valid `SEARCH` expression as first argument. You can set the boost
factor for that sub-expression via the second parameter. Documents that match
boosted parts of the search expression will get higher scores.

### Dataset

[IMDB movie dataset](example-datasets.md#imdb-movie-dataset)

### View definition

#### `search-alias` View

```js
db.imdb_vertices.ensureIndex({
  name: "inv-text",
  type: "inverted",
  fields: [
    { name: "description", analyzer: "text_en" }
  ]
});

db._createView("imdb_alias", "search-alias", { indexes: [
  { collection: "imdb_vertices", index: "inv-text" }
] });
```

#### `arangosearch` View

```json
{
  "links": {
    "imdb_vertices": {
      "fields": {
        "description": {
          "analyzers": [
            "text_en"
          ]
        }
      }
    }
  }
}
```

### AQL queries

Prefer `galaxy` over the other keywords:

_`search-alias` View:_

```aql
FOR doc IN imdb_alias
  SEARCH doc.description IN TOKENS("amazing action world alien sci-fi science documental", "text_en")
      OR BOOST(doc.description IN TOKENS("galaxy", "text_en"), 5)
  SORT BM25(doc) DESC
  LIMIT 10
  RETURN {
    title: doc.title,
    description: doc.description,
    score: BM25(doc)
  }
```

_`arangosearch` View:_

```aql
FOR doc IN imdb
  SEARCH ANALYZER(doc.description IN TOKENS("amazing action world alien sci-fi science documental", "text_en")
      OR BOOST(doc.description IN TOKENS("galaxy", "text_en"), 5), "text_en")
  SORT BM25(doc) DESC
  LIMIT 10
  RETURN {
    title: doc.title,
    description: doc.description,
    score: BM25(doc)
  }
```

| title | description | score |
|:------|:------------|:------|
| Star Trek Collection | Star Trek a futuristic **science** fiction franchise. … **galaxies** to explore, and cool skin tight suits to beam up in … | 64.87849426269531 |
| Alien Tracker | In a **galaxy** far away, **alien** criminals organize a spectacular prison break. … Cole is the **Alien** Tracker … | 63.959991455078125 |
| Stitch! The Movie | … the **galaxy's** most wanted extraterrestrial … Dr. Jumba brought one of his **alien** "experiments" to Hawaii. | 63.39030075073242 |
| The Hitchhiker's Guide to the Galaxy | Mere seconds before the Earth is to be demolished by an **alien** construction crew … a new edition of "The Hitchhiker's Guide to the **Galaxy**." | 63.37282943725586 |
| Stargate: The Ark of Truth | … it may be in the Ori's own home **galaxy**. … SG-1 travels to the Ori **galaxy** … in a distant **galaxy** fighting two powerful enemies. | 61.784141540527344 |
| The Ice Pirates | … the most precious commodity in the **galaxy** is water. … unreachable centre of the **galaxy** … The **galaxy** is ruled by an evil emperor … | 61.78216552734375 |
| Star Wars: Episode III: Revenge of the Sith | … leading a massive clone army into a **galaxy**-wide battle against the Separatists. … to rule the **galaxy**, the Republic crumbles … | 59.79429244995117 |
| Star Wars: Episode II - Attack of the Clones | … not only has the **galaxy** undergone significant change, but so have Obi-Wan Kenobi, Padmé Amidala, and Anakin Skywalker … | 55.723636627197266 |
| Macross Plus | … a new aircraft (Shinsei Industries' YF-19 & General **Galaxy's** YF-21) for Project Super Nova, to choose the newest successor to the VF-11 | 55.722259521484375 |
| Star Trek | The fate of the **galaxy** rests in the hands of bitter rivals. One, James Kirk, is a delinquent, thrill-seeking Iowa farm boy. The other, Spock, a Vulcan, … | 55.717037200927734 |

If you are an information retrieval expert and want to fine-tuning the
weighting schemes at query time, then you can do so. The `BM25()` function
accepts free coefficients as parameters to turn it into BM15 for instance:

_`search-alias` View:_

```aql
FOR doc IN imdb_alias
  SEARCH doc.description IN TOKENS("amazing action world alien sci-fi science documental", "text_en")
      OR BOOST(doc.description IN TOKENS("galaxy", "text_en"), 5)
  LET score = BM25(doc, 1.2, 0)
  SORT score DESC
  LIMIT 10
  RETURN {
    title: doc.title,
    description: doc.description,
    score
  }
```

_`arangosearch` View:_

```aql
FOR doc IN imdb
  SEARCH ANALYZER(doc.description IN TOKENS("amazing action world alien sci-fi science documental", "text_en")
      OR BOOST(doc.description IN TOKENS("galaxy", "text_en"), 5), "text_en")
  LET score = BM25(doc, 1.2, 0)
  SORT score DESC
  LIMIT 10
  RETURN {
    title: doc.title,
    description: doc.description,
    score
  }
```

| title | description | score |
|:------|:------------|:------|
| Stargate: The Ark of Truth | … it may be in the Ori's own home **galaxy**. … SG-1 travels to the Ori **galaxy** … in a distant **galaxy** fighting two powerful enemies. | 42.88237380981445 |
| The Ice Pirates | … the most precious commodity in the **galaxy** is water. … unreachable centre of the **galaxy** … The **galaxy** is ruled by an evil emperor … | 42.88237380981445 |
| Star Wars: Episode III: Revenge of the Sith | … leading a massive clone army into a **galaxy**-wide battle against the Separatists. … to rule the **galaxy**, the Republic crumbles … | 39.27024841308594 |
| Alien Tracker | In a **galaxy** far away, **alien** criminals organize a spectacular prison break. … Cole is the **Alien** Tracker … | 38.43224334716797 |
| Star Trek Collection | Star Trek a futuristic **science** fiction franchise. … **galaxies** to explore, and cool skin tight suits to beam up in … | 38.42367935180664 |
| Stitch! The Movie | … the **galaxy's** most wanted extraterrestrial … Dr. Jumba brought one of his **alien** "experiments" to Hawaii. | 37.563819885253906 |
| The Hitchhiker's Guide to the Galaxy | Mere seconds before the Earth is to be demolished by an **alien** construction crew … a new edition of "The Hitchhiker's Guide to the **Galaxy**." | 37.563819885253906 |
| Critters 4 | … he gets a message that it would be illegal to extinguish the race from the **galaxy**. … | 32.99643325805664
| Alien Agent | A lawman from another **galaxy** must stop an invading force from building a gateway to planet Earth. | 32.99643325805664
| Star Trek | The fate of the **galaxy** rests in the hands of bitter rivals. One, James Kirk, is a delinquent, thrill-seeking Iowa farm boy. The other, Spock, a Vulcan, … | 32.99643325805664 |

You can also calculate a custom score, taking into account additional fields
of the document.

Match movies with the (normalized) phrase `star war` in the title and calculate
a custom score based on BM25 and the movie runtime to favor longer movies:

```aql
FOR doc IN imdb
/* `search-alias` View:
FOR doc IN imdb_alias
*/
  SEARCH PHRASE(doc.title, "Star Wars", "text_en")
  LET score = BM25(doc) * LOG(doc.runtime + 1)
  SORT score DESC
  RETURN {
    title: doc.title,
    runtime: doc.runtime,
    bm25: BM25(doc),
    score
  }
```

| title | runtime | bm25 | score |
|:------|:--------|:-----|:------|
| **Star Wars**: Episode II - Attack of the Clones | 142 | 16.900253295898438 | 83.87333131958185 |
| **Star Wars**: Episode III: Revenge of the Sith | 140 | 16.900253295898438 | 83.63529564797363 |
| **Star Wars**: Episode VI - Return of the Jedi | 135 | 16.900253295898438 | 83.02511192427228 |
| **Star Wars**: Episode I - The Phantom Menace | 133 | 16.81275749206543 | 82.34619279156092 |
| **Star Wars**: Episode V: The Empire Strikes Back | 124 | 16.900253295898438 | 81.59972515247492 |
| **Star Wars**: Episode IV - A New Hope | 121 | 16.81275749206543 | 80.76884081187906 |
| The **Star Wars** Holiday Special | 97 | 16.569408416748047 | 75.97019873160025 |
| **Star Wars**: The Clone Wars | 90 | 16.569408416748047 | 74.74227347404823 |
| **Star Wars**: Revelations | 47 | 16.13064956665039 | 62.44498690901793 |
| **Star Wars** Collection | null | 16.13064956665039 | 0 |
