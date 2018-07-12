# Endpoints

## Table of Contents:

- [Food Endpoints](#food-endpoints)
  - [GET /api/v1/foods](#foods-index)
  - [GET /api/v1/foods/:id](#food-show)
  - [PUTS /api/v1/foods](#create-food)
  - [PATCH /api/v1/foods/:id](#update-food)
  - [DELETE /api/v1/foods/:id](#delete-food)
- [Meal Endpoints](#meal-endpoints)
  - [GET /api/v1/meals](#meals-index)
  - [GET /api/v1/meals/:meal_id/foods](#meal-show)
  - [PUTS /api/v1/meals/:meal_id/foods/:id](#add-food-to-meal)
  - [DELETE /api/v1/foods/:id](#remove-food-from-meal)

### Food Endpoints:

#### Foods Index

Endpoint: '/api/v1/foods' <br>
Verb: GET

Returns a JSON list of all foods that have been created and their attributes.

Example response:
```json
[{
  "id": 1,
  "name": "Banana",
  "calories": 150
},
{
  "id": 2,
  "name": "Pizza",
  "calories": 1500
}]
```

#### Food Show

Endpoint: '/api/v1/foods/:id' <br>
Verb: GET

Returns the food object associated with the id parameter.

Example response:
```json
{
  "id": 1,
  "name": "Banana",
  "calories": 150
}
```

#### Create Food

Endpoint: '/api/v1/foods' <br>
Verb: PUTS

Pass food parameters in the body of the post request to create a new food item.

Example request body:
```json
{
  "food": {
    "name": "Name of food here",
    "calories": "Calories here"
    }
}
```

Example response:
```json
[{
  "id": 1,
  "name": "Banana",
  "calories": 150
},
{
  "id": 2,
  "name": "Pizza",
  "calories": 1500
}]
```

#### Update Food

Endpoint: '/api/v1/foods/:id' <br>
Verb: PATCH

Updates the food object associated with the id parameter.

Example response:
```json
{
  "id": 1,
  "name": "Updated name",
  "calories": 150
}
```

#### Delete Food

Endpoint: '/api/v1/foods/:id' <br>
Verb: DELETE

Deletes the food object associated with the id parameter.

---

### Meal Endpoints:

#### Meals Index

Endpoint: '/api/v1/meals' <br>
Verb: GET

Returns a list of all meals, and the foods that are associated with each meal.

Example response:
```json
[{
  "id": 1,
  "name": "Breakfast",
  "foods": [
            {
              "id": 1,
              "name": "Banana",
              "calories": 150
            },
            {
              "id": 6,
              "name": "Yogurt",
              "calories": 550
            },
            {
              "id": 12,
              "name": "Apple",
              "calories": 220
            }
          ]
}]
```

#### Meal Show

Endpoint: '/api/v1/meals/:meal_id/foods' <br>
Verb: GET

Returns the meal identified by :meal_id and all of its associated foods.

Example response:
```json
{
  "id": 1,
  "name": "Breakfast",
  "foods": [
            {
              "id": 1,
              "name": "Banana",
              "calories": 150
            },
            {
              "id": 6,
              "name": "Yogurt",
              "calories": 550
            }
          ]
}
```

#### Add Food to Meal

Endpoint: '/api/v1/meals/:meal_id/foods/:id' <br>
Verb: PUTS

Create a relationship between the food identified by the :id parameter and the meal identified by the :meal_id parameter.

Example response:
```json
{
  "message": "Successfully added FOODNAME to MEALNAME"
}
```

#### Remove Food from Meal

Endpoint: '/api/v1/meals/:meal_id/foods/:id' <br>
Verb: DELETE

Destroys the relationship between the food identified by the :id parameter and the meal identified by the :meal_id parameter.

Example response:
```json
{
  "message": "Successfully removed FOODNAME from MEALNAME"
}
```
