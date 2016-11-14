<!DOCTYPE html>

<html>
    <head>
        <link rel="stylesheet" type="text/css" href="/static/css/grid.css" />
        <link rel="stylesheet" type="text/css" href="/static/css/ext-theme-crisp-all-debug.css" />
        <link rel="stylesheet" type="text/css" href="/static/css/sencha-charts-all-debug.css" />

        <!--query builder css -->
        <link rel="stylesheet" type="text/css" href="/static/css/querybuilder/bootstrap.min.css">
        <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-multiselect.css">
        <link rel="stylesheet" type="text/css" href="/static/css/querybuilder/bootstrap-select.css">
        <link rel="stylesheet" type="text/css" href="/static/css/querybuilder/bootstrap-datepicker3.css">
        <link rel="stylesheet" type="text/css" href="/static/css/querybuilder/awesome-bootstrap-checkbox.css">
        <link rel="stylesheet" type="text/css" href="/static/css/querybuilder/query-builder.default.css">
        <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css">

        <!-- GC -->
        <script type="text/javascript" src="/static/js/querybuilder/jquery.js"></script>
        <script type="text/javascript" src="/static/js/querybuilder/bootstrap.min.js"></script>
        <script type="text/javascript" src="/static/js/querybuilder/bootstrap-select.js"></script>
        <script type="text/javascript" src="/static/js/querybuilder/bootbox.js"></script>
        <script type="text/javascript" src="/static/js/querybuilder/bootstrap-datepicker.js"></script>
        <!-- sql plug in -->
        <script type="text/javascript" src="/static/js/querybuilder/sql-parser.js"></script>
        <script type="text/javascript" src="/static/js/querybuilder/doT.js"></script>
        <script type="text/javascript" src="/static/js/querybuilder/jQuery.extendext.js"></script>
        <script type="text/javascript" src="/static/js/querybuilder/query-builder.js"></script>
        <script type="text/javascript" src="/static/js/bootstrap-multiselect.js"></script>
        <script type="text/javascript" src="/static/js/moment.js"></script>
        <script type="text/javascript" src="/static/js/daterangepicker.js"></script>

        <script type="text/javascript" src="/static/js/ext-all.js"></script>

        <!-- page specific   -->
        <script type="text/javascript" src="/static/js/binding.js"></script>
        <script type="text/javascript" src="/static/js/mul_Select.js"></script>
        <script type="text/javascript" src="/static/js/custRangeDatePicker.js"></script>

        <style>
            .home{
                font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
                font-size: 1rem;
                line-height: 1.5;
                color: #555555;
                background-color: #c1e2b3;
                height: 100%;
            }
        </style>

    </head>

    <body class="home">
    <div id="dateRangePicker" class="pull-right" style="background: #fff; cursor: pointer; padding: 5px 10px; border: 1px solid #ccc; width: 18%">
        <i class="glyphicon glyphicon-calendar fa fa-calendar"></i>&nbsp;
        <span></span> <b class="caret"></b>
    </div>
    <div class="btn-group pull-right group-actions" id="divSelectGroup">
        <select id="sParentCategory" multiple="multiple"></select>
        <select id="sCategory" multiple="multiple"></select>
        <select id="sCourseName" multiple="multiple"></select>
    </div>

        <table width="100%">
            <tr>
                <td></td>
                <td></td>
            </tr>
            <tr>
                <td colspan="3"><div id="query-builder"></div></td>
            </tr>
            <tr>
                <td colspan="3"><div id="import-grid"></div></td>
            </tr>
        </table>
        <button id="btnPopWin" class="btn btn-primary btn-lg" style="display: none" data-toggle="modal" data-target="#myModal">
        </button>
        <div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <button type="button" class="close" data-dismiss="modal" aria-hidden="true">
                            &times;
                        </button>
                        <h4 class="modal-title" id="myModalLabel">
                            选择文件类型下载
                        </h4>
                    </div>
                    <div class="modal-body">
                        <table>
                            <tr>
                                <td> <div class="radio radio-primary">
                                    <input type="radio" name="radio1" id="rdExcel" value="xlsx" checked>
                                    <label for="rdExcel">
                                        Excel
                                    </label>
                                </div></td>
                                <td width="20px"></td>
                                <td><div class="radio radio-primary">
                                    <input type="radio" name="radio1" id="rdCSV" value="csv">
                                    <label for="rdCSV">
                                        CSV
                                    </label>
                                </div></td>
                            </tr>
                        </table>
                    </div>
                    <div class="modal-footer">
                        <button type="button" id="btnClose" class="btn btn-default" data-dismiss="modal">关闭
                        </button>
                        <button id="btnConfirm" type="button" class="btn btn-primary">
                            确认
                        </button>
                    </div>
                </div><!-- /.modal-content -->
            </div><!-- /.modal -->
        </div>
    </body>

</html>
<script>
    window.BootData = {
        Title : {{.ImportDataDefinition.GridTitle}},
        Columns : {{.ImportDataDefinition.Columns}},
        Filters : {{.ImportDataDefinition.Filters}},
        Fields : {{.ImportDataDefinition.Fields}},
        SelectGroup : {{.ImportDataDefinition.SelectGroup}},
    }

</script>

<script>
    $('#query-builder').queryBuilder({
        plugins: ['bt-tooltip-errors','bt-selectpicker','bt-checkbox'],
        filters: JSON.parse(window.BootData.Filters)
    });

    $("#btnConfirm").click(function (e) {
        var extensions = $("#rdExcel")[0].checked ? $("#rdExcel")[0].value : $("#rdCSV")[0].value;
        $("#btnClose").click();
        $.ajax({
            type: "POST",
            url: '/api/importData',
            data: { "extensions":extensions, "filters":filters, "cols" : JSON.stringify(visibleCols) },
            success: function(result) {
                var resultData = JSON.parse(result)
                if(resultData.successful == true)
                    window.location.href = "/static/tmpFile/mbData."    + extensions;
            }
        });
    })

    var start,end,filters;
    var visibleCols = [];

    $(".rules-group-header").append($('#dateRangePicker'))
    $(".rules-group-header").append($('#divSelectGroup'))

</script>

