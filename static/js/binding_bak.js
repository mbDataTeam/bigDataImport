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
        fields: [
            'Id',
            'Title',
            'Author',
            'Manufacturer',
            'Product',
        ]
    });

    // create the Data Store
    var store = Ext.create('Ext.data.Store', {
        model: 'ImportData',
        proxy: {
            // load using HTTP
            type: 'ajax',
            url: '/api/fetchData',
            // the return will be XML, so lets set up a reader
            reader: {
                type: 'json',
                //record: 'Item',
                totalProperty  : 'total'
            }
        }
    });

    // create the grid
    var grid = Ext.create('Ext.grid.Panel', {
        bufferedRenderer: false,
        store: store,
        columns: [
            {text: "Id", width: 120, dataIndex: 'Id', sortable: true},
            {text: "Author", width: 120, dataIndex: 'Author', sortable: true},
            {text: "Title", flex: 1, dataIndex: 'Title', sortable: true},
            {text: "Manufacturer", width: 125, dataIndex: 'Manufacturer', sortable: true},
            {text: "Product", width: 125, dataIndex: 'Product', sortable: true}
        ],
        forceFit: true,
        height:400,
        split: true,
        region: 'north'
    });

    Ext.create('Ext.Panel', {
        renderTo: 'import-grid',
        frame: true,
        title: 'Book List',
        width: 580,
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

