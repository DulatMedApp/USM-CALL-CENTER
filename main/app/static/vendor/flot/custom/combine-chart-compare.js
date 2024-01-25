$(function () {    
 var data24Hours = [
        [0, 0], [1, 0], [2, 0], [3, 0], [4, 0], [5, 0], [6, 5], [7, 25], [8, 38], [9, 5], [10, 3], [11, 15], [12,40],
        [13, 50], [14, 20], [15, 15], [16, 10], [17, 7], [18, 10], [19, 0], [20, 0], [21, 0], [22, 0], [23, 0]
    ];



 var ticks = [
        [0, "00:00"], [1, "01:00"], [2, "02:00"], [3, "03:00"], [4, "04:00"], [5, "05:00"],
        [6, "06:00"], [7, "07:00"], [8, "08:00"], [9, "09:00"], [10, "10:00"], [11, "11:00"],
        [12, "12:00"], [13, "13:00"], [14, "14:00"], [15, "15:00"], [16, "16:00"], [17, "17:00"],
        [18, "18:00"], [19, "19:00"], [20, "20:00"], [21, "21:00"], [22, "22:00"], [23, "23:00"]
    ];

	var data = [{
		label: "Последние 24 часа",
		data: data24Hours,
		lines: {
			show: true, lineWidth: 2
		},
		points: {
			show:true,
			radius: 4,
			fill: true,
			fillColor: "#ffffff",
			lineWidth: 3,
			symbol: "circle"
		}
	}];

	var options = {
		series: {
			shadowSize: 0,
			bars: {
				lineWidth: 2,
			}
		},
		xaxis: {
			ticks: ticks
		},
		grid: {
			hoverable: true,
			clickable: false,
			borderWidth: 1,
			tickColor: '#f5f6fa',
			borderColor: '#f5f6fa',
		},
		legend:{
			show: true,
			position: 'se',
			noColumns: 0, //In single row
			// labelBoxBorderColor: "#000000",
			// container: $("#legendcontainer"),
		},
		tooltip: true,
		tooltipOpts: {
			content: '%x: %y'
		},
		colors: ['#c776db', '#d599e4', '#f1ddf6'],
	};
	var plot = $.plot($("#combineChartCompare"), 
	data
, options);  


   // Добавляем текстовые метки к точкам
    for (var i = 0; i < data24Hours.length; i++) {
        var point = data24Hours[i];
        var label = point[1].toString(); // Преобразуем значение в строку
        var offset = plot.pointOffset({ x: point[0], y: point[1] });
        $("<div class='data-point-label'>" + label + "</div>")
            .css({
                position: 'absolute',
                left: offset.left - 10, // Настройте смещение по горизонтали
                top: offset.top - 20,  // Настройте смещение по вертикали
                display: 'none',
                color: '#333',
                'font-size': '12px',
                'font-weight': 'bold'
            })
            .appendTo(plot.getPlaceholder())
            .fadeIn('slow');
    }
});