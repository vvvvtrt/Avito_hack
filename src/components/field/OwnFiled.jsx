import React from "react";
import SetButton from "../../buttons/setButtons";
import Graph from "../graphView/graph";
import CatGraph from "../graphView/graph-cat";
import './field.css'

function Field(){


    return(<div className="field">
        <div className="wrapper-graph">
            <Graph/>
            <CatGraph/>
            <SetButton/>

        </div>

    </div>)
}

export default Field;