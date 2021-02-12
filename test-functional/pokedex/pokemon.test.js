import * as api from './functions/api_pokemon'

const testFile = 'Pokemon'

describe('pokedex-server', () => {
    beforeAll(async () => {

    })
    afterAll(async () => {

    })
    test('Filter Pokemon by type', async () => {
        const expectedResult = {
            data: {
                listPokemon:
                    {
                      count: 1,
                      edges:  [
                      {
                       id: "046",
                           isFavorite: false,
                           name: "Paras",
                           number: 0,
                           types:  [
                             "Bug", "Grass",
                               ],
                           weight:  {
                             maximum: "6.08kg",
                                 minimum: "4.73kg",
                               },
                     },
               ],
               limit: 1, offset: "046",
                    }
            }
        }
        const result = await api.ListPokemonByType({
            searchType: "Bug",
            limit: 1
        }, testFile)
        expect(result.data).toEqual(expectedResult)
    })
    test('Search for Pokemon by text.', async () => {
        const expectedResult = {
            data: {
                listPokemon:
                    {
                        count: 1,
                        edges:  [
                            {
                                id: "046",
                                isFavorite: false,
                                name: "Paras",
                                number: 0,
                                types:  [
                                    "Bug", "Grass",
                                ],
                                weight:  {
                                    maximum: "6.08kg",
                                    minimum: "4.73kg",
                                },
                            },
                        ],
                        limit: 1, offset: "046",
                    }
            }
        }

        const result = await api.ListPokemonByName({
            searchName: "Paras",
            limit: 1
        }, testFile)
        expect(result.data).toEqual(expectedResult)
    })
    test('Add and remove a Pokemon to and from your Favorites', async () => {
        const expectedResult = {
            data: {
                updateFavoritePokemon: {
                    id: "004",
                    number: 0,
                    name: "Charmander",
                    classification: "LizardPokémon",
                    isFavorite: true,
                    types: [
                        "Fire"
                    ],
                    weaknesses: [
                        "Water",
                        "Ground",
                        "Rock"
                    ],
                    weight: {
                        minimum: "7.44kg",
                        maximum: "9.56kg"
                    }
                }
            }
        }

        const result = await api.UpdateFavoritePokemon({
            isfavorite: true,
            id: "004"
        }, testFile)
        expect(result.data).toEqual(expectedResult)
    })

    test('Retrieve All Pokemons and Favorite Pokemons', async () => {
        const expectedResult = {
            data: {
                listPokemon:
                    {
                        count: 1,
                        edges:  [
                            {
                                id: "004",
                                isFavorite: true,
                                name: "Charmander",
                                number: 0,
                                types:  [
                                    "Fire",
                                ],
                                weight:  {
                                    maximum: "9.56kg",
                                    minimum: "7.44kg",
                                },
                            },
                        ],
                        limit: 1, offset: "004",
                    }
            }
        }


        const result = await api.ListPokemon({
            isfavorite: true,
            limit: 1
        }, testFile)
        expect(result.data).toEqual(expectedResult)
    })

    test('Get Pokemon details.', async () => {
        const expectedResult = {
            data: {
                getPokemonById: {
                    id: "004",
                    number: 0,
                    name: "Charmander",
                    classification: "LizardPokémon",
                    isFavorite: true,
                    types: [
                        "Fire"
                    ],
                    weaknesses: [
                        "Water",
                        "Ground",
                        "Rock"
                    ],
                    weight: {
                        minimum: "7.44kg",
                        maximum: "9.56kg"
                    }

                }
            }
        }
        const result = await api.GetPokemonById({
            id: "004",
        }, testFile)
        expect(result.data).toEqual(expectedResult)
    })

})
