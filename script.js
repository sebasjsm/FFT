// Leer el JSON
const jsonData = [{"fase":0,"mag":28},{"fase":1.9634954,"mag":10.452504},{"fase":2.3561945,"mag":5.656854},{"fase":2.7488935,"mag":4.329569},{"fase":3.1415927,"mag":4},{"fase":-2.7488935,"mag":4.329569},{"fase":-2.3561945,"mag":5.656854},{"fase":-1.9634954,"mag":10.452504}];

// Obtener los valores de fase y magnitud del JSON
const fases = jsonData.map(d => d.fase);
const magnitudes = jsonData.map(d => d.mag);

// Configurar el gráfico
const ctx = document.getElementById('myChart').getContext('2d');
const chart = new Chart(ctx, {
    type: 'line',
    data: {
        labels: fases, // Eje x
        datasets: [{
            label: 'Magnitud',
            data: magnitudes, // Eje y
            backgroundColor: 'rgba(255, 99, 132, 0.2)',
            borderColor: 'rgba(255, 99, 132, 1)',
            borderWidth: 1
        }]
    },
    options: {
        scales: {
            yAxes: [{
                beginAtZero: true,
                ticks: {
                    callback: function(value, index, values) {
                        return value.toFixed(4); // Ajusta la cantidad de decimales
                    }
                }
            }]
        }
    }
});
