---
title: Geo-spatial functions in AQL
menuTitle: Geo
weight: 35
description: >-
  The following helper functions can use geo indexes, but do not have to in all cases
archetype: default
---
## Geo utility functions

The following helper functions **can** use geo indexes, but do not have to in
all cases. You can use all of these functions in combination with each other,
and if you have configured a geo index it may be utilized,
see [Geo Indexing](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md).

### DISTANCE()

`DISTANCE(latitude1, longitude1, latitude2, longitude2) → distance`

Calculate the distance between two arbitrary points in meters (as birds
would fly). The value is computed using the haversine formula, which is based
on a spherical Earth model. It's fast to compute and is accurate to around 0.3%,
which is sufficient for most use cases such as location-aware services.

- **latitude1** (number): the latitude of the first point
- **longitude1** (number): the longitude of the first point
- **latitude2** (number): the latitude of the second point
- **longitude2** (number): the longitude of the second point
- returns **distance** (number): the distance between both points in **meters**

```aql
// Distance from Brandenburg Gate (Berlin) to ArangoDB headquarters (Cologne)
DISTANCE(52.5163, 13.3777, 50.9322, 6.94) // 476918.89688380965 (~477km)

// Sort a small number of documents based on distance to Central Park (New York)
FOR doc IN coll // e.g. documents returned by a traversal
  SORT DISTANCE(doc.latitude, doc.longitude, 40.78, -73.97)
  RETURN doc
```

### GEO_CONTAINS()

`GEO_CONTAINS(geoJsonA, geoJsonB) → bool`

Checks whether the [GeoJSON object](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md#geojson) `geoJsonA`
fully contains `geoJsonB` (every point in B is also in A). The object `geoJsonA`
has to be of type _Polygon_ or _MultiPolygon_. For other types containment is
not well-defined because of numerical stability problems.

- **geoJsonA** (object): first GeoJSON object
- **geoJsonB** (object): second GeoJSON object, or a coordinate array in
  `[longitude, latitude]` order
- returns **bool** (bool): `true` if every point in B is also contained in A,
  otherwise `false`

{{< info >}}
ArangoDB follows and exposes the same behavior as the underlying
S2 geometry library. As stated in the S2 documentation:

> Point containment is defined such that if the sphere is subdivided
> into faces (loops), every point is contained by exactly one face.
> This implies that linear rings do not necessarily contain their vertices.

As a consequence, a linear ring or polygon does not necessarily contain its
boundary edges!
{{< /info >}}

You can optimize queries that contain a `FILTER` expression of the following
form with an S2-based [geospatial index](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md):

```aql
FOR doc IN coll
  FILTER GEO_CONTAINS(geoJson, doc.geo)
  ...
```

In this example, you would create the index for the collection `coll`, on the
attribute `geo`. You need to set the `geoJson` index option to `true`.
The `geoJson` variable needs to evaluate to a valid GeoJSON object. Also note
the argument order: the stored document attribute `doc.geo` is passed as the
second argument. Passing it as the first argument, like
`FILTER GEO_CONTAINS(doc.geo, geoJson)` to test whether `doc.geo` contains
`geoJson`, cannot utilize the index.

### GEO_DISTANCE()

`GEO_DISTANCE(geoJsonA, geoJsonB, ellipsoid) → distance`

Return the distance between two GeoJSON objects in meters, measured from the
**centroid** of each shape. For a list of supported types see the
[geo index page](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md#geojson).

- **geoJsonA** (object): first GeoJSON object, or a coordinate array in
  `[longitude, latitude]` order
- **geoJsonB** (object): second GeoJSON object, or a coordinate array in
  `[longitude, latitude]` order
- **ellipsoid** (string, *optional*): reference ellipsoid to use.
  Supported are `"sphere"` (default) and `"wgs84"`.
- returns **distance** (number): the distance between the centroid points of
  the two objects on the reference ellipsoid in **meters**

```aql
LET polygon = {
  type: "Polygon",
  coordinates: [[[-11.5, 23.5], [-10.5, 26.1], [-11.2, 27.1], [-11.5, 23.5]]]
}
FOR doc IN collectionName
  LET distance = GEO_DISTANCE(doc.geometry, polygon) // calculates the distance
  RETURN distance
```

You can optimize queries that contain a `FILTER` expression of the following
form with an S2-based [geospatial index](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md):

```aql
FOR doc IN coll
  FILTER GEO_DISTANCE(geoJson, doc.geo) <= limit
  ...
```

In this example, you would create the index for the collection `coll`, on the
attribute `geo`. You need to set the `geoJson` index option to `true`.
`geoJson` needs to evaluate to a valid GeoJSON object. `limit` must be a
distance in meters; it cannot be an expression. An upper bound with `<`,
a lower bound with `>` or `>=`, or both, are equally supported.

You can also optimize queries that use a `SORT` condition of the following form
with a geospatial index:

```aql
  SORT GEO_DISTANCE(geoJson, doc.geo)
```

The index covers returning matches from closest to furthest away, or vice versa.
You may combine such a `SORT` with a `FILTER` expression that utilizes the
geospatial index, too, via the [`GEO_DISTANCE()`](#geo_distance),
[`GEO_CONTAINS()`](#geo_contains), and [`GEO_INTERSECTS()`](#geo_intersects)
functions.

### GEO_AREA()

<small>Introduced in: v3.5.1</small>

`GEO_AREA(geoJson, ellipsoid) → area`

Return the area for a polygon or multi-polygon on a sphere with the
average Earth radius, or an ellipsoid. For a list of supported types
see the [geo index page](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md#geojson).

- **geoJson** (object): a GeoJSON object
- **ellipsoid** (string, *optional*): reference ellipsoid to use.
  Supported are `"sphere"` (default) and `"wgs84"`.
- returns **area** (number): the area of the polygon in **square meters**

```aql
LET polygon = {
  type: "Polygon",
  coordinates: [[[-11.5, 23.5], [-10.5, 26.1], [-11.2, 27.1], [-11.5, 23.5]]]
}
RETURN GEO_AREA(polygon, "wgs84")
```

### GEO_EQUALS()

`GEO_EQUALS(geoJsonA, geoJsonB) → bool`

Checks whether two GeoJSON objects are equal or not. For a list of supported
types see the [geo index page](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md#geojson).

- **geoJsonA** (object): first GeoJSON object.
- **geoJsonB** (object): second GeoJSON object.
- returns **bool** (bool): `true` if they are equal, otherwise `false`.

```aql
LET polygonA = GEO_POLYGON([
  [-11.5, 23.5], [-10.5, 26.1], [-11.2, 27.1], [-11.5, 23.5]
])
LET polygonB = GEO_POLYGON([
  [-11.5, 23.5], [-10.5, 26.1], [-11.2, 27.1], [-11.5, 23.5]
])
RETURN GEO_EQUALS(polygonA, polygonB) // true
```

```aql
LET polygonA = GEO_POLYGON([
  [-11.1, 24.0], [-10.5, 26.1], [-11.2, 27.1], [-11.1, 24.0]
])
LET polygonB = GEO_POLYGON([
  [-11.5, 23.5], [-10.5, 26.1], [-11.2, 27.1], [-11.5, 23.5]
])
RETURN GEO_EQUALS(polygonA, polygonB) // false
```

### GEO_INTERSECTS()

`GEO_INTERSECTS(geoJsonA, geoJsonB) → bool`

Checks whether the [GeoJSON object](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md#geojson) `geoJsonA`
intersects with `geoJsonB` (i.e. at least one point in B is also in A or vice-versa).

- **geoJsonA** (object): first GeoJSON object
- **geoJsonB** (object): second GeoJSON object, or a coordinate array in
  `[longitude, latitude]` order
- returns **bool** (bool): true if B intersects A, false otherwise

You can optimize queries that contain a `FILTER` expression of the following
form with an S2-based [geospatial index](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md):

```aql
FOR doc IN coll
  FILTER GEO_INTERSECTS(geoJson, doc.geo)
  ...
```

In this example, you would create the index for the collection `coll`, on the
attribute `geo`. You need to set the `geoJson` index option to `true`.
`geoJson` needs to evaluate to a valid GeoJSON object. Also note
the argument order: the stored document attribute `doc.geo` is passed as the
second argument. Passing it as the first argument, like
`FILTER GEO_INTERSECTS(doc.geo, geoJson)` to test whether `doc.geo` intersects
`geoJson`, cannot utilize the index.

### GEO_IN_RANGE()

<small>Introduced in: v3.8.0</small>

`GEO_IN_RANGE(geoJsonA, geoJsonB, low, high, includeLow, includeHigh) → bool`

Checks whether the distance between two [GeoJSON objects](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md#geojson)
lies within a given interval. The distance is measured from the **centroid** of
each shape.

- **geoJsonA** (object\|array): first GeoJSON object, or a coordinate array
  in `[longitude, latitude]` order
- **geoJsonB** (object\|array): second GeoJSON object, or a coordinate array
  in `[longitude, latitude]` order
- **low** (number): minimum value of the desired range
- **high** (number): maximum value of the desired range
- **includeLow** (bool, optional): whether the minimum value shall be included
  in the range (left-closed interval) or not (left-open interval). The default
  value is `true`
- **includeHigh** (bool): whether the maximum value shall be included in the
  range (right-closed interval) or not (right-open interval). The default value
  is `true`
- returns **bool** (bool): whether the evaluated distance lies within the range

### IS_IN_POLYGON()

Determine whether a point is inside a polygon.

{{< warning >}}
The `IS_IN_POLYGON()` AQL function is **deprecated** as of ArangoDB 3.4.0 in
favor of the new [`GEO_CONTAINS()` AQL function](#geo_contains), which works with
[GeoJSON](https://tools.ietf.org/html/rfc7946) Polygons and MultiPolygons.
{{< /warning >}}

`IS_IN_POLYGON(polygon, latitude, longitude) → bool`

- **polygon** (array): an array of arrays with 2 elements each, representing the
  points of the polygon in the format `[latitude, longitude]`
- **latitude** (number): the latitude of the point to search
- **longitude** (number): the longitude of the point to search
- returns **bool** (bool): `true` if the point (`[latitude, longitude]`) is
  inside the `polygon` or `false` if it's not. The result is undefined (can be
  `true` or `false`) if the specified point is exactly on a boundary of the
  polygon.

```aql
// checks if the point (latitude 4, longitude 7) is contained inside the polygon
IS_IN_POLYGON( [ [ 0, 0 ], [ 0, 10 ], [ 10, 10 ], [ 10, 0 ] ], 4, 7 )
```

---

`IS_IN_POLYGON(polygon, coord, useLonLat) → bool`

The 2nd parameter can alternatively be specified as an array with two values.

By default, each array element in `polygon` is expected to be in the format
`[latitude, longitude]`. This can be changed by setting the 3rd parameter to `true` to
interpret the points as `[longitude, latitude]`. `coord` is then also interpreted in
the same way.

- **polygon** (array): an array of arrays with 2 elements each, representing the
  points of the polygon
- **coord** (array): the point to search as a numeric array with two elements
- **useLonLat** (bool, *optional*): if set to `true`, the coordinates in
  `polygon` and the coordinate pair `coord` are interpreted as
  `[longitude, latitude]` (like in GeoJSON). The default is `false` and the
  format `[latitude, longitude]` is expected.
- returns **bool** (bool): `true` if the point `coord` is inside the `polygon`
  or `false` if it's not. The result is undefined (can be `true` or `false`) if
  the specified point is exactly on a boundary of the polygon.

```aql
// checks if the point (lat 4, lon 7) is contained inside the polygon
IS_IN_POLYGON( [ [ 0, 0 ], [ 0, 10 ], [ 10, 10 ], [ 10, 0 ] ], [ 4, 7 ] )

// checks if the point (lat 4, lon 7) is contained inside the polygon
IS_IN_POLYGON( [ [ 0, 0 ], [ 10, 0 ], [ 10, 10 ], [ 0, 10 ] ], [ 7, 4 ], true )
```

## GeoJSON Constructors

The following helper functions are available to easily create valid GeoJSON
output. In all cases you can write equivalent JSON yourself, but these functions
will help you to make all your AQL queries shorter and easier to read.

### GEO_LINESTRING()

`GEO_LINESTRING(points) → geoJson`

Construct a GeoJSON LineString.
Needs at least two longitude/latitude pairs.

- **points** (array): an array of `[longitude, latitude]` pairs
- returns **geoJson** (object): a valid GeoJSON LineString

```aql
---
name: aqlGeoLineString_1
description: ''
---
RETURN GEO_LINESTRING([
    [35, 10], [45, 45]
])
```

### GEO_MULTILINESTRING()

`GEO_MULTILINESTRING(points) → geoJson`

Construct a GeoJSON MultiLineString.
Needs at least two elements consisting valid LineStrings coordinate arrays.

- **points** (array): array of LineStrings
- returns **geoJson** (object): a valid GeoJSON MultiLineString

```aql
---
name: aqlGeoMultiLineString_1
description: ''
---
RETURN GEO_MULTILINESTRING([
    [[100.0, 0.0], [101.0, 1.0]],
    [[102.0, 2.0], [101.0, 2.3]]
])
```

### GEO_MULTIPOINT()

`GEO_MULTIPOINT(points) → geoJson`

Construct a GeoJSON LineString. Needs at least two longitude/latitude pairs.

- **points** (array): an array of `[longitude, latitude]` pairs
- returns **geoJson** (object): a valid GeoJSON Point

```aql
---
name: aqlGeoMultiPoint_1
description: ''
---
RETURN GEO_MULTIPOINT([
    [35, 10], [45, 45]
])
```

### GEO_POINT()

`GEO_POINT(longitude, latitude) → geoJson`

Construct a valid GeoJSON Point.

- **longitude** (number): the longitude portion of the point
- **latitude** (number): the latitude portion of the point
- returns **geoJson** (object): a GeoJSON Point

```aql
---
name: aqlGeoPoint_1
description: ''
---
RETURN GEO_POINT(1.0, 2.0)
```

### GEO_POLYGON()

`GEO_POLYGON(points) → geoJson`

Construct a GeoJSON Polygon. Needs at least one array representing
a linear ring. Each linear ring consists of an array with at least four
longitude/latitude pairs. The first linear ring must be the outermost, while
any subsequent linear ring will be interpreted as holes.

For details about the rules, see [GeoJSON polygons](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md#polygon).

- **points** (array): an array of (arrays of) `[longitude, latitude]` pairs
- returns **geoJson** (object\|null): a valid GeoJSON Polygon

A validation step is performed using the S2 geometry library. If the
validation is not successful, an AQL warning is issued and `null` is
returned.

Simple Polygon:

```aql
---
name: aqlGeoPolygon_1
description: ''
---
RETURN GEO_POLYGON([
    [0.0, 0.0], [7.5, 2.5], [0.0, 5.0], [0.0, 0.0]
])
```

Advanced Polygon with a hole inside:

```aql
---
name: aqlGeoPolygon_2
description: ''
---
RETURN GEO_POLYGON([
    [[35, 10], [45, 45], [15, 40], [10, 20], [35, 10]],
    [[20, 30], [30, 20], [35, 35], [20, 30]]
])
```

### GEO_MULTIPOLYGON()

`GEO_MULTIPOLYGON(polygons) → geoJson`

Construct a GeoJSON MultiPolygon. Needs at least two Polygons inside.
See [GEO_POLYGON()](#geo_polygon) and [GeoJSON MultiPolygons](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md#multipolygon) for the rules of Polygon and MultiPolygon construction.

- **polygons** (array): an array of arrays of arrays of `[longitude, latitude]` pairs
- returns **geoJson** (object\|null): a valid GeoJSON MultiPolygon

A validation step is performed using the S2 geometry library, if the
validation is not successful, an AQL warning is issued and `null` is
returned.

MultiPolygon comprised of a simple Polygon and a Polygon with hole:

```aql
---
name: aqlGeoMultiPolygon_1
description: ''
---
RETURN GEO_MULTIPOLYGON([
    [
        [[40, 40], [20, 45], [45, 30], [40, 40]]
    ],
    [
        [[20, 35], [10, 30], [10, 10], [30, 5], [45, 20], [20, 35]],
        [[30, 20], [20, 15], [20, 25], [30, 20]]
    ]
])
```

## Geo Index Functions

{{< warning >}}
The AQL functions `NEAR()`, `WITHIN()` and `WITHIN_RECTANGLE()` are
deprecated starting from version 3.4.0.
Please use the [Geo utility functions](#geo-utility-functions) instead.
{{< /warning >}}

AQL offers the following functions to filter data based on
[geo indexes](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md). These functions require the collection
to have at least one geo index. If no geo index can be found, calling this
function will fail with an error at runtime. There is no error when explaining
the query however.

### NEAR()

{{< warning >}}
`NEAR()` is a deprecated AQL function from version 3.4.0 on.
Use [`DISTANCE()`](#distance) in a query like this instead:

```aql
FOR doc IN coll
  SORT DISTANCE(doc.latitude, doc.longitude, paramLatitude, paramLongitude) ASC
  RETURN doc
```
Assuming there exists a geo-type index on `latitude` and `longitude`, the
optimizer will recognize it and accelerate the query.
{{< /warning >}}

`NEAR(coll, latitude, longitude, limit, distanceName) → docArray`

Return at most *limit* documents from collection *coll* that are near
*latitude* and *longitude*. The result contains at most *limit* documents,
returned sorted by distance, with closest distances being returned first.
Optionally, the distances in meters between the specified coordinate pair
(*latitude* and *longitude*) and the stored coordinate pairs can be returned as
well. To make use of that, the desired attribute  name for the distance result
has to be specified in the *distanceName* argument. The result documents will
contain the distance value in an attribute of that name.

- **coll** (collection): a collection
- **latitude** (number): the latitude of the point to search
- **longitude** (number): the longitude of the point to search
- **limit** (number, *optional*): cap the result to at most this number of
  documents. The default is 100. If more documents than *limit* are found,
  it is undefined which ones will be returned.
- **distanceName** (string, *optional*): include the distance (in meters)
  between the reference point and the stored point in the result, using the
  attribute name *distanceName*
- returns **docArray** (array): an array of documents, sorted by distance
  (shortest distance first)

### WITHIN()

{{< warning >}}
`WITHIN()` is a deprecated AQL function from version 3.4.0 on.
Use [`DISTANCE()`](#distance) in a query like this instead:

```aql
FOR doc IN coll
  LET d = DISTANCE(doc.latitude, doc.longitude, paramLatitude, paramLongitude)
  FILTER d <= radius
  SORT d ASC
  RETURN doc
```

Assuming there exists a geo-type index on `latitude` and `longitude`, the
optimizer will recognize it and accelerate the query.
{{< /warning >}}

`WITHIN(coll, latitude, longitude, radius, distanceName) → docArray`

Return all documents from collection *coll* that are within a radius of *radius*
around the specified coordinate pair (*latitude* and *longitude*). The documents
returned are sorted by distance to the reference point, with the closest
distances being returned first. Optionally, the distance (in meters) between the
reference point and the stored point can be returned as well. To make
use of that, an attribute name for the distance result has to be specified in
the *distanceName* argument. The result documents will contain the distance
value in an attribute of that name.

- **coll** (collection): a collection
- **latitude** (number): the latitude of the point to search
- **longitude** (number): the longitude of the point to search
- **radius** (number): radius in meters
- **distanceName** (string, *optional*): include the distance (in meters)
  between the reference point and stored point in the result, using the
  attribute name *distanceName*
- returns **docArray** (array): an array of documents, sorted by distance
  (shortest distance first)

### WITHIN_RECTANGLE()

{{< warning >}}
`WITHIN_RECTANGLE()` is a deprecated AQL function from version 3.4.0 on. Use
[`GEO_CONTAINS()`](#geo_contains) and a GeoJSON polygon instead - but note that
this uses geodesic lines from version 3.10.0 onward
(see [GeoJSON interpretation](../../index-and-search/indexing/working-with-indexes/geo-spatial-indexes.md#geojson-interpretation)):

```aql
LET rect = GEO_POLYGON([ [
  [longitude1, latitude1], // bottom-left
  [longitude2, latitude1], // bottom-right
  [longitude2, latitude2], // top-right
  [longitude1, latitude2], // top-left
  [longitude1, latitude1], // bottom-left
] ])
FOR doc IN coll
  FILTER GEO_CONTAINS(rect, [doc.longitude, doc.latitude])
  RETURN doc
```

Assuming there exists a geo-type index on `latitude` and `longitude`, the
optimizer will recognize it and accelerate the query.
{{< /warning >}}

`WITHIN_RECTANGLE(coll, latitude1, longitude1, latitude2, longitude2) → docArray`

Return all documents from collection *coll* that are positioned inside the
bounding rectangle with the points (*latitude1*, *longitude1*) and (*latitude2*,
*longitude2*). There is no guaranteed order in which the documents are returned.

- **coll** (collection): a collection
- **latitude1** (number): the latitude of the bottom-left point to search
- **longitude1** (number): the longitude of the bottom-left point to search
- **latitude2** (number): the latitude of the top-right point to search
- **longitude2** (number): the longitude of the top-right point to search
- returns **docArray** (array): an array of documents, in random order
