<!DOCTYPE html>

<html>
    <head>
        <link rel="stylesheet" type="text/css" href="/static/css/grid.css" />
        <link rel="stylesheet" type="text/css" href="/static/css/ext-theme-crisp-all-debug.css" />
        <link rel="stylesheet" type="text/css" href="/static/css/sencha-charts-all-debug.css" />

        <!--query builder css -->
        <link rel="stylesheet" type="text/css" href="/static/css/querybuilder/bootstrap.min.css">
        <link rel="stylesheet" type="text/css" href="/static/css/querybuilder/bootstrap-select.css">
        <link rel="stylesheet" type="text/css" href="/static/css/querybuilder/bootstrap-datepicker3.css">
        <link rel="stylesheet" type="text/css" href="/static/css/querybuilder/awesome-bootstrap-checkbox.css">
        <link rel="stylesheet" type="text/css" href="/static/css/querybuilder/query-builder.default.css">

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

        <script type="text/javascript" src="/static/js/ext-all.js"></script>

        <!-- page specific   -->
        <script type="text/javascript" src="/static/js/binding.js"></script>

    </head>

    <body>
        <div id="query-builder"></div>
        <div id="import-grid"></div>
    </body>

</html>
<script>
    window.BootData = {
        Title : {{.ImportDataDefinition.GridTitle}},
        Columns : {{.ImportDataDefinition.Columns}},
        Filters : {{.ImportDataDefinition.Filters}},
        Fields : {{.ImportDataDefinition.Fields}},
        TableName : {{.ImportDataDefinition.TableName}},
    }

</script>

<script>
    $('#query-builder').queryBuilder({
        plugins: ['bt-tooltip-errors','bt-selectpicker','bt-checkbox'],
        filters: JSON.parse(window.BootData.Filters)
    });

    // hide theme select control
    // $("#ext-element-6").css("display","none")
    // $("#options-toolbar").css("display","none")
</script>
