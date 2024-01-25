var chart10 = c3.generate({
	bindto: '#donutChart',
	data: {
		columns: [
			['Пропущенные', 5200],
			['Входящие', 9709],
			['Исходящие', 2200],
		],
		type : 'donut',
		colors: {
			Пропущенные: '#FF6576',
			Входящие: '#76D12C',
			Исходящие: '#50B0F5',
		},
		onclick: function (d, i) { console.log("onclick", d, i); },
		onmouseover: function (d, i) { console.log("onmouseover", d, i); },
		onmouseout: function (d, i) { console.log("onmouseout", d, i); }
	},
	 donut: {
    title: "Всего",
		 width: 40,
		expand: true, // развернуть радиус
  },
  onrendered: function () {
    d3.select('#donutChart .c3-chart-arcs-title')
      .append('tspan')
      .attr('x', 0)
      .attr('dy', 15)
			.text("11 922")
		 .style('font-weight', 'bold');
  },
});