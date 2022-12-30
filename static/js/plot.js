function drawPlot(x, y, title, ytitle) {
    const trace1 = {
        x: x,
        y: y,
        type: 'scatter'
    };

    const layout = {
        autosize: true,
        title: {
            text: title,
            font: {
                size: 24
            },
            xref: 'paper',
            x: 0.05,
        },
        y: 0.5,
        traceorder: 'reversed',
        font: {size: 16},
        xaxis: {
            title: 'Time'
        },
        yaxis: {
            title: ytitle
        }
    };

    const data = [trace1];

    const config = {responsive: true};

    Plotly.newPlot('myDiv', data, layout, config);
}