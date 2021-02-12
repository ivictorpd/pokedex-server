import { SendRequest } from '../../request'

const apiTestFile = 'api_pokemon'

export async function ListPokemon(data, testFile) {
  data = `
        query listPokemonFavorite{
            listPokemon(input:{filter:{isFavorite:${data.isfavorite}},limit:${data.limit}}){
              edges{
                id
                number
                name
                weight{
                  minimum
                  maximum
                }
                types
                isFavorite
              }
              offset
              count
              limit
            }
          }
    `
  return SendRequest(data, apiTestFile, testFile)
}

export async function ListPokemonByType(data, testFile) {
  data = `
        query ListPokemonByType{
          listPokemon(input:{filter:{searchType:"${data.searchType}"},limit:${data.limit}}){
              edges{
                id
                number
                name
                weight{
                  minimum
                  maximum
                }
                types
                isFavorite
              }
              offset
              count
              limit
            }
          }
    `
  return SendRequest(data, apiTestFile, testFile)
}

export async function ListPokemonByName(data, testFile) {
  data = `
        query ListPokemonByName{
          listPokemon(input:{filter:{searchName:"${data.searchName}"},limit:${data.limit}}){
              edges{
                id
                number
                name
                weight{
                  minimum
                  maximum
                }
                types
                isFavorite
              }
              offset
              count
              limit
            }
          }
    `
  return SendRequest(data, apiTestFile, testFile)
}

export async function GetPokemonById(data, testFile) {
  data = `
        query GetPokemonById{
            getPokemonById(id:"${data.id}"){
             id
              number
              name
              classification
              types
              isFavorite
              weaknesses
              classification
              weight{
                minimum
                maximum
              }
            }
          }
    `
  return SendRequest(data, apiTestFile, testFile)
}

export async function UpdateFavoritePokemon(data, testFile) {
  data = `
        mutation updateFavoritePokemon{
            updateFavoritePokemon(id:"${data.id}", isFavorite: ${data.isfavorite}){
               id
                number
                name
                classification
                isFavorite
                types
                weaknesses
                weight{
                  minimum
                  maximum
                }
              }
            }
    `
  return SendRequest(data, apiTestFile, testFile)
}

export async function ListPokemonTypes(data, testFile) {
  data = `
       query listPokemonTypes{
            listPokemonTypes
          }
    `
  return SendRequest(data, apiTestFile, testFile)
}