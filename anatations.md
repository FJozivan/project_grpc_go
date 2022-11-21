## Criando aquivos proto buffers
* Criar arquivo.proto
* Definir syntax proto que sera usada exemplo 
    ```go
    syntax = "proto3";
    ```
* Definir pacote do go que sera usado para ser acessado pelo go para busca das estruturas
    ```go
    package pb; // Define o pacote
    option go_package = "internal/pb"; // Define o caminho do pacote
    ```
* Definir estruturas
    ```go
    message Category { // O "message" Define o inicio de uma nova estrutura
        string id = 1; // Os atributos são definidos pelo: tipo + nome + posição do mesmo na estrutura
        string name = 2;
        string description = 3;
    }
    ```
* È possível definir novas estruturas com estruturas ja criadas
    ```go
        message  CategoryResponse{
            Category category = 1;
        }
    ```
* Dentro do arquivo é possível definir as funções dos serviço que farão parte da interface a ser gerada ja se utilizando das estruturas para saber parâmetros de entradas e os retornos.
    ```go
        service CategoryService{ // Define o nome do serviço
            rpc CreateCategory(CategoryCreateResquest) returns (Category){};// Define-se as funções genéricas que que fao parte da interface do serviço gerada.
        }
    ```