# QuickCalcGo

QuickCalcGo es una calculadora CLI desarrollada en Go que permite realizar operaciones matemáticas básicas.

## Instalación de Go en Ubuntu

Para utilizar QuickCalcGo, necesitas tener Go instalado en tu sistema. Aquí te explicamos cómo instalar la última versión de Go en Ubuntu:

1.  **Actualiza tu sistema**:

    ```bash
    sudo apt update
    sudo apt upgrade
    ```

2.  **Descarga el archivo tar de Go**:
    Visita [la página oficial de Go](https://golang.org/dl/) para encontrar la versión más reciente y reemplaza el URL en el siguiente comando:

    ```bash
    wget https://golang.org/dl/go1.17.2.linux-amd64.tar.gz
    ```

    Asegúrate de reemplazar `1.x.x` con la versión actual de Go.

3.  **Extrae el archivo tar**:

    ```bash
    sudo tar -C /usr/local -xzf go1.17.2.linux-amd64.tar.gz
    ```

    Recuerda ajustar la versión de Go en el nombre del archivo.

4.  **Configura las Variables de Entorno**:
    Edita tu archivo `.bashrc` o `.profile`:

    ```bash
    nano ~/.bashrc
    ```

    Añade las siguientes líneas al final del archivo:

    ```bash
    export GOPATH=$HOME/go
    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
    ```

    Guarda y cierra el archivo. Luego, aplica los cambios:

    ```bash
    source ~/.bashrc
    ```

5.  **Verifica la Instalación**:
    Ejecuta `go version` para verificar que Go se haya instalado correctamente.

## Ejecución del Proyecto QuickCalcGo

Una vez instalado Go, puedes clonar y ejecutar QuickCalcGo siguiendo estos pasos:

1.  **Clona el Repositorio** (ajusta el comando si el proyecto está en un repositorio en línea):

    ```bash
    git clone git@github.com:Daniel-Santiago-Acosta-1013/QuickCalcGo.git
    cd QuickCalcGo
    ```

2.  **Ejecuta el Proyecto**:
    Desde el directorio raíz del proyecto, ejecuta:
    ```bash
    go run ./cmd
    ```

Ahora puedes comenzar a usar QuickCalcGo para realizar cálculos matemáticos en la línea de comandos.

## Contribuciones

Las contribuciones a QuickCalcGo son bienvenidas. Si tienes alguna sugerencia o mejora, no dudes en abrir un issue o un pull request en el repositorio.
