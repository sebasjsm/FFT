
// Obtén los datos en formato JSON desde el servidor
//fase eje x y magnitud eje y
const data = [
    {"mag": 28, "fase": 0},
    {"mag": 10.452504, "fase": 1.963495},
    {"mag": 5.656854, "fase": 2.356194},
    {"mag": 4.329569, "fase": 2.748893},
    {"mag": 4.000000, "fase": 3.141593},
    {"mag": 4.329569, "fase": -2.748893},
    {"mag": 5.656854, "fase": -2.356194},
    {"mag": 10.452504, "fase": -1.963495}
];

// Prepara los datos para el gráfico

const magnitudes = data.map(d => d.mag);;
const fases = data.map(d => d.fase);

// Configura el gráfico
const ctx = document.getElementById('myChart').getContext('2d');
const chart = new Chart(ctx, {
    type: 'line',
    data: {
        labels: magnitudes, //eje y
        datasets: [{
            label: 'Magnitud',
            data: fases, //eje x
            backgroundColor: 'rgba(255, 99, 132, 0.2)',
            borderColor: 'rgba(255, 99, 132, 1)',
            borderWidth: 1
        }, //{
            //label: 'Fase',
            //data: fases,
            //backgroundColor: 'rgba(255, 99, 132, 0.2)',
            //borderColor: 'rgba(255, 99, 132, 1)',
            //borderWidth: 1
        //}
    ]
    },
    options: {
        scales: {
            yAxes:[ {
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
