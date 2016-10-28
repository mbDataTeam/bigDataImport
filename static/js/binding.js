/**
 * Created by Bill Hu on 2016/10/13.
 */
Ext.require([
    'Ext.grid.*',
    'Ext.data.*',
    'Ext.panel.*',
    'Ext.layout.container.Border'
]);

Ext.onReady(function(){
    Ext.define('ImportData',{
        extend: 'Ext.data.Model',
        fields: window.BootData.Fields,
    });

    // create the Data Store
    var store = Ext.create('Ext.data.Store', {
        model: 'ImportData',
        //pageSize: 15,
        id: 'gridStore',
        proxy: {
            // load using HTTP
            type: 'ajax',
            url: '/api/fetchData',
            reader: {
                type: 'json',
                rootProperty: 'data',
                totalProperty  : 'rowCount'
            }
        }
    });

    // create the grid
    var grid = Ext.create('Ext.grid.Panel', {
        //bufferedRenderer: false,
        store: store,
        columns: JSON.parse(window.BootData.Columns),
        bbar: {
            type: 'pagingtoolbar',
            store: store,
            displayInfo: true,
            displayMsg: 'Displaying {0} to {1} of {2} &nbsp;records ',
            emptyMsg: "No records to display&nbsp;"
        },
        viewConfig:{
            autoFill:true
        },
        height: 360,
        split: true,
        region: 'north'
    });

    Ext.create('Ext.Panel', {
        renderTo: 'import-grid',
        frame: true,
        title: window.BootData.Title,
        header:{
            items:[
                {
                    xtype: 'button',
                    text: 'Search',
                    handler: function () {
                        var filters =""
                        if(window.BootData.SelectGroup) {
                            var firstCataValues = getCheckedParameter($("#sParentCategory"));
                            if(firstCataValues == "" || firstCataValues == null){
                                filters = ""
                            }
                            else {
                                filters = " Parent_Category_Id in (" + firstCataValues+ ")"
                                var secCataValues = getCheckedParameter($("#sCategory"));
                                secCataValues == "" ? filters : filters = filters + " and Category_Id in (" + secCataValues + ")"
                                var thirdCataValues = getCheckedParameter($("#sCourseName"));
                                thirdCataValues == "" ? filters : filters = filters +" and course_id in (" + thirdCataValues + ")";
                            }
                        }

                        var root = $('#query-builder').queryBuilder('getModel');
                        rules = root.model.root.rules;
                        if (rules.length == 0 ||(rules.length == 1 && !rules[0].filter)){
                            grid.store.clearData();
                            grid.view.refresh();
                            store.load(
                                {
                                    params : {
                                        filters: filters
                                    }
                                }
                            )
                        }
                        else {
                            var result = $('#query-builder').queryBuilder('getSQL', false);
                            if (result.sql.length > 0) {
                                grid.store.clearData();
                                grid.view.refresh();
                                filters == "" ? (filters = result.sql) : (filters = filters + " and " + result.sql)
                                bootbox.alert({
                                    title: "sql语句",
                                    message: "<P>" + filters + "</P>"
                                });
                                store.load(
                                    {
                                        params : {
                                            filters: filters
                                        }
                                    }
                                )
                            }

                        }

                    }
                },
                {
                    xtype: 'button',
                    text: 'Export',
                    handler: function () {
                        if(Ext.getStore("gridStore").totalCount > 0) {
                            $("#btnPopWin").click()
                        }

                        /*
                        var result = $('#query-builder').queryBuilder('getSQL', false);
                        if (result.sql.length > 0) {
                            result.sql = "select * from "+window.BootData.TableName+" where " + result.sql;
                            bootbox.alert({
                                title: "sql语句",
                                message: '<pre class="code-popup">' + result.sql + '</pre>'
                            });
                        }*/
                    }
                }
            ]
        },
        width: '100%',
        height: 400,
        layout: 'border',
        items: [
            grid, {
                id: 'detailPanel',
                region: 'center',
                bodyPadding: 7,
                bodyStyle: "background: #ffffff;"
            }]
    });

    store.load();
});

