$(function () {

    var ds = [], data = [], chartOptions, textLabels = [];

    ds.push([[2100, 1], [3300, 2], [3900, 3], [4500, 4], [5200, 5], [3100, 6]]);
    ds.push([[1500, 1], [2200, 2], [2900, 3], [2300, 4], [3800, 5], [1600, 6]]);
    ds.push([[600, 1], [1300, 2], [1900, 3], [3500, 4], [2700, 5], [3200, 6]]);

    var departmentNames = ds[0].map(function (item) { return item[0]; });

    for (var i = 0; i < ds.length; i++) {
        data.push({
            label: ds[i][0][1],
            data: ds[i],
            bars: {
                horizontal: true,
                barWidth: .20,
                show: true,
                fill: true,
                lineWidth: 0,
                order: i,
                fillColor: { colors: [{ opacity: 1 }, { opacity: 1 }] }
            },
            datalabels: {
                show: true,
                align: 'center',
                color: '#333',
                font: {
                    size: 12,
                    family: 'Arial, sans-serif',
                    weight: 'bold'
                },
                formatter: function (value, series) {
                    return departmentNames[series.datapoints.points[0] / 2 - 1];
                }
            }
        });

        // Add text labels
        textLabels.push({
            text: departmentNames[i],
            x: ds[i][0][0],
            y: ds[i][0][1] + 0.2,  // Adjust the vertical position of the text label
            align: 'left',
            textAlign: 'center',
            font: {
                size: 12,
                family: 'Arial, sans-serif',
                weight: 'bold',
                color: '#333'
            },
            color: '#333'
        });
    }

    chartOptions = {
        xaxis: {},
        yaxis: {
            ticks: ds[0].map(function (item) { return [item[1], item[0]]; }),
            tickLength: 0,
            tickFormatter: function (value, axis) {
                for (var i = 0; i < ds[0].length; i++) {
                    if (ds[0][i][1] == value) {
                        return ds[0][i][0];
                    }
                }
            },
            font: {
                size: 12,
                family: 'Arial, sans-serif',
                weight: 'bold',
                color: '#333'
            }
        },
        grid: {
            hoverable: true,
            clickable: false,
            borderWidth: 1,
            tickColor: '#f5f6fa',
            borderColor: '#f5f6fa',
        },
        series: {
            stack: true
        },
        legend: {
            position: 'se'
        },
        bars: {
            horizontal: true,
            show: true,
            barWidth: .50,
            fill: true,
            lineWidth: 0,
            fillColor: { colors: [{ opacity: 1 }, { opacity: 1 }] }
        },
        shadowSize: 0,
        tooltip: true,

        tooltipOpts: {
            content: '%s: %x'
        },
        colors: ['#ff7c4c', '#ffa27f', '#ffc7b2'],
    };

    var holder = $('#stacked-horizontal-chart');
    if (holder.length) {
        $.plot(holder, data, chartOptions);
        for (var i = 0; i < textLabels.length; i++) {
            holder.find('.flot-text').eq(i).css('color', textLabels[i].color);
            holder.find('.flot-text').eq(i).text(textLabels[i].text);
            holder.find('.flot-text').eq(i).attr('x', chartOptions.xaxis.p2c(textLabels[i].x));
            holder.find('.flot-text').eq(i).attr('y', chartOptions.yaxis.p2c(textLabels[i].y));
            holder.find('.flot-text').eq(i).attr('align', textLabels[i].align);
            holder.find('.flot-text').eq(i).attr('text-anchor', textLabels[i].textAlign);
            holder.find('.flot-text').eq(i).attr('font-family', textLabels[i].font.family);
            holder.find('.flot-text').eq(i).attr('font-size', textLabels[i].font.size);
            holder.find('.flot-text').eq(i).attr('font-weight', textLabels[i].font.weight);
        }
    }

});
