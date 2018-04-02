Keen.ready(function(){
  window.Common.daterange_change(showGraphs);

    // Expenses (pie)
  var expenses_pie = new Keen.Dataviz()
    .el('#expenses')
    .type('pie')
    .height(280)
    .title('Expenses')
    .prepare();

  // Profit time line by client
  var profit = new Keen.Dataviz()
    .el('#profit')
    .type('area')
    .height(280)
    .stacked(false)
    .title('Profit')
    .prepare();

  function showGraphs(start, end){
        $.ajax({
            type: "POST",
            url: "/api/v1/expenses",
            data: JSON.stringify({Directive: "list", Params:[start, end]}),
            success: function(data) {
                if (data.hasOwnProperty("reason") || data.hasOwnProperty("code")) {
                    expenses_pie.message("Could not request expenses data");
                    console.error(data);
                    return;
                }
                console.info("data", data);
                expenses_pie
                    .data(data)
                    .sortGroups('desc')
                    .render();
            },
            dataType: "json",
            error: function(e) {
                console.error(e);
                expenses_pie.message("Could not request expenses data");
            }
        });

      $.ajax({
          type: "POST",
          url: "/api/v1/profit",
          data: JSON.stringify({Directive: "list", Params:[start, end]}),
          success: function(data) {
              if (data.hasOwnProperty("reason") || data.hasOwnProperty("code")) {
                  profit.message("Could not request expenses data");
                  console.error(data);
                  return;
              }
              console.info("data", data);
              profit
                  .data(data)
                  .sortGroups('desc')
                  .render();
          },
          dataType: "json",
          error: function(e) {
              console.error(e);
              profit.message("Could not request expenses data");
          }
      });
    }
    showGraphs(window.Common.DEFAULT_START, window.Common.DEFAULT_END);
});
