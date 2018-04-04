Keen.ready(function(){
    // Expenses (pie)
  var expenses_pie = new Keen.Dataviz()
    .el('#pie-chart')
    .type('pie')
    .height(window.Common.Height)
    .title('Expenses')
    .prepare();

  // Profit
  var profit = new Keen.Dataviz()
    .el('#bar-chart')
    .type('bar')
    .height(window.Common.Height)
    .stacked(false)
    .title('Profit')
    .prepare();

  function showGraphs(start, end){
        $.ajax({
            type: "GET",
            url: "/api/v1/expenses",
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

      function showTable(data, $container) {
          $container.html("<div class='keen-dataviz-title'>"+data.Name+"</div>");
          var $table = $("<table class='table'></table>");
          $container.append($table);
          function addRow(row, isHeader) {
              var $tr = $("<tr></tr>");
              $table.append($tr);
              row.forEach(function (elem) {
                  var td="<td>"+elem+"</td>";
                  if (isHeader) {
                      td="<th>"+elem+"</th>";
                  }
                  $tr.append(td);
              });
          }
          if (data.hasOwnProperty("Header") && data.Header) {
              addRow(data.Header, true);
          }
          data.Data.forEach(function (row) {
              addRow(row, false);
          });
      }

      var $qualityAndReleases = $("#quality-and-releases");
      $.ajax({
          type: "POST",
          url: "/api/v1/quality-and-releases",
          data: JSON.stringify({Directive: "list", Params:[start, end]}),
          success: function(data) {
              if (data.hasOwnProperty("reason") || data.hasOwnProperty("code")) {
                  $qualityAndReleases.text("Could not request quality and releases data");
                  console.error(data);
                  return;
              }
              console.info("data", JSON.stringify(data));
              showTable(data, $qualityAndReleases);
          },
          dataType: "json",
          error: function(e) {
              console.error(e);
              alert("Could not request expenses data");
          }
      });

      var $otherKey = $("#other-key");
      $.ajax({
          type: "POST",
          url: "/api/v1/other-key",
          data: JSON.stringify({Directive: "list", Params:[start, end]}),
          success: function(data) {
              if (data.hasOwnProperty("reason") || data.hasOwnProperty("code")) {
                  $otherKey.text("Could not request quality and releases data");
                  console.error(data);
                  return;
              }
              console.info("data", JSON.stringify(data));
              showTable(data, $otherKey);
          },
          dataType: "json",
          error: function(e) {
              console.error(e);
              expenses_pie.message("Could not request expenses data");
          }
      });
    }
    showGraphs(window.Common.Start, window.Common.End);
});
