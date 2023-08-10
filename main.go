package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Cargar productos desde el archivo CSV
	products, err := loadProducts("products.csv")
	if err != nil {
		fmt.Println("Error al cargar los productos:", err)
		return
	}

	// Mapa para almacenar las cantidades compradas por código de producto
	quantities := make(map[string]int)

	// Mapa para almacenar los subtotales por código de producto
	subtotals := make(map[string]float64)

	// Proceso de compra
	scanner := bufio.NewScanner(os.Stdin)
	total := 0.0

	fmt.Println("Bienvenido al punto de venta!")

	for {
		fmt.Print("Ingrese el código de producto (000 para finalizar): ")
		scanner.Scan()
		code := scanner.Text()

		if code == "000" {
			break
		}

		// Verificar si el código de producto es válido
		product, exists := products[code]
		if !exists {
			fmt.Println("Código de producto inválido")
			continue
		}

		fmt.Print("Ingrese la cantidad: ")
		scanner.Scan()
		quantityStr := scanner.Text()

		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			fmt.Println("Cantidad inválida")
			continue
		}

		subtotal := float64(quantity) * product.price
		subtotals[code] += subtotal
		total += subtotal
		quantities[code] += quantity

		fmt.Printf("Producto: %s\nCantidad: %d\nSubtotal: %.2f\n\n", product.name, quantity, subtotal)
	}

	// Imprimir el ticket
	fmt.Println("\nTicket:")
	fmt.Println("----------------------------------------")
	for code, quantity := range quantities {
		product := products[code]
		subtotal := subtotals[code]
		fmt.Printf("%s\t%s\t%d\t%.2f\t%.2f\n", code, product.name, quantity, product.price, subtotal)
	}
	fmt.Println("----------------------------------------")
	fmt.Printf("Total del ticket: %.2f\n", total)
	fmt.Println("Gracias por su compra!")
}

type Product struct {
	code  string
	name  string
	price float64
}

func loadProducts(filename string) (map[string]Product, error) {
	products := make(map[string]Product)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		code := line[0]
		name := line[1]
		price, _ := strconv.ParseFloat(line[2], 64)

		products[code] = Product{code, name, price}
	}

	return products, nil
}
