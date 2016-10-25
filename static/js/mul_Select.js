/**
 * Created by Bill Hu on 2016/10/24.
 */

$(document).ready(function () {
    var pCatalog = $("#sParentCategory");
    configSelect(pCatalog, "Parent Catalog list","pCatalog");

    var catalog = $("#sCategory");
    configSelect(catalog, "Catalog list","catalog");

    var course = $("#sCourseName");
    configSelect(course, "Course list");

    fillFirstSelectData(window.BootData.CompanyId, "pCatalog");

});


// get data to fill select control
function fillFirstSelectData(ids, selectType ) {

    $.ajax({
        url: "/api/fillSelect",
        data: { "compIds": ids , "selectType" : selectType},
        type: "POST",
        dataType: "json",
        success: function (data) {
            //alert("(Get) [Data from Server] " + JSON.stringify(data));
            if (data != null) {
                loadMultiSelectOption(data, $("#sParentCategory"), true);
                var idList = getIdList(data);
                //alert(idList);
                fillSecondData(idList)
            }
        }
    });
}

function fillSecondData(ids) {
    $.ajax({
        url: "/api/fillSelect",
        data: { "ids": ids , "compIds": window.BootData.CompanyId,"selectType" : "catalog"},
        type: "POST",
        dataType: "json",
        success: function (data) {
            //alert("(Get) [Data from Server] " + JSON.stringify(data));
            if (data != null) {
                loadMultiSelectGroupOption(data,$("#sCategory"), true);
                var idList = getIdList(data);
                fillThirdData(idList)
            }
        }
    });
}

function fillThirdData(ids) {
    $.ajax({
        url: "/api/fillSelect",
        data: { "ids": ids , "compIds": window.BootData.CompanyId, "selectType" : "course"},
        type: "POST",
        dataType: "json",
        success: function (data) {
            //alert("(Get) [Data from Server] " + JSON.stringify(data));
            if (data != null) {
                loadMultiSelectGroupOption(data, $("#sCourseName"), true);
            }
        }
    });
}

function loadMultiSelectOption(jsonData, $menu, checked) {
    $menu.children("option").each(function (index, option) {
        $(option).remove();
    });

    var count = jsonData.length;
    for(var i = 0; i< count; i++){
        $menu.append($('<option>', {
            value: jsonData[i].OptionValue,
            text:  jsonData[i].OptionText
        }));
    }

    if (checked) {
        $menu.children('option').prop('selected', true);
    }
    $menu.multiselect('rebuild');

}

/*
 schema define
 <optgroup label="Group 1" class="group-1">
 <option value="1-1">Option 1.1</option>
 <option value="1-2" selected="selected">Option 1.2</option>
 </optgroup>
 <optgroup label="Group 2" class="group-2">
 <option value="2-1">Option 2.1</option>
 <option value="2-2">Option 2.2</option>
 </optgroup>
 */
function loadMultiSelectGroupOption(jsonData,$menu, checked) {
    $menu.children("optgroup").each(function (index, optgroup) {
        $(optgroup).children('option').remove();
        $(optgroup).remove();
    });

    var count = jsonData.length;
    var groups = [];
    for(var i = 0; i< count; i++){
        var hasGroupName = checkGroupExsit(groups,jsonData[i].GroupName)
        if(!hasGroupName){
            groups.push(jsonData[i].GroupName);
            var group = $("<optgroup label="+jsonData[i].GroupName+"></optgroup>");
            $(group).append($('<option>', {
                value: jsonData[i].OptionValue,
                text:  jsonData[i].OptionText
            }));
            $menu.append(group);
        }else{
                $menu.children("optgroup").each(function (index, optgroup) {
                if(optgroup.label == jsonData[i].GroupName ) {
                    $(optgroup).append($('<option>', {
                        value: jsonData[i].OptionValue,
                        text:  jsonData[i].OptionText
                    }));
                }
            });
        }
    }
    if (checked) {
        $menu.children("optgroup").each(function (index, optgroup) {
            $(optgroup).children('option').prop('selected', true);
        });
    }
    $menu.multiselect('rebuild');
    groups.length = 0;
}

function checkGroupExsit(groups,groupName) {
    var hasGroup = false;
    $.each(groups,function (i,n) {
        if(groups[i] == groupName){
            hasGroup =true
        }
    })
    return hasGroup
}

function configSelect($selectObj,txt,selectType) {
    $selectObj.multiselect({
        nonSelectedText: txt,
        buttonWidth: "240",
        maxHeight: "200",
        includeSelectAllOption: true,
        enableFiltering: true,
        enableCaseInsensitiveFiltering: false,
        enableClickableOptGroups: true,
        //disableIfEmpty: true,
        buttonText: function (options, select) {
            return txt + "(" + options.length + ")";
        },
        onDropdownShow: function(event) {
            $(".multiselect-container label").removeClass("checkbox"); //show checkbox
        },

        onChange: function(option, checked, select) {
            var values = getCheckedParameter($selectObj);
            if(selectType == "pCatalog"){
                fillSecondData(values);
            }
            else if(selectType == "catalog"){
                fillThirdData(values);
            }
        },

        onSelectAll: function () {
            var values = getCheckedParameter($selectObj);
            if(selectType == "pCatalog"){
                fillSecondData(values);
            }
            else if(selectType == "catalog"){
                fillThirdData(values);
            }
        },
        onDeselectAll: function() {
            var values = getCheckedParameter($selectObj);
            if(selectType == "pCatalog"){
                fillSecondData(values);
            }
            else if(selectType == "catalog"){
                fillThirdData(values);
            }
        }
    });
}

function getCheckedParameter($tag) {
    var p = "";
    $tag.find("option:selected").each(function () {
        p += $(this).val() + ",";
    })
    p = p.substring(0, p.length - 1);
    return p;
}

//format like : id1,id2,id3,...
function getIdList(data) {
    var strIds =""
    for (var i=0; i< data.length; i++){
        strIds += data[i].OptionValue + ",";
    }
    if(strIds.length > 0){
        strIds = strIds.substring(0,strIds.length-1);
    }
    return strIds;
}

