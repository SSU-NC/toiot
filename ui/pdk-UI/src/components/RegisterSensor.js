import React, { Component } from 'react';
// form : https://getbootstrap.com/docs/4.0/components/forms/?
// add, delete input : https://codesandbox.io/s/00xq32n3pn?from-embed=&file=/src/index.js

class RegisterSensor extends Component {
    constructor(props) {
        super(props);
        this.state = {
            value_list: [{value_name:""}],
            name: ""
        };
        // this.handleNameChange = this.handleNameChange.bind(this);
        // this.handleAddClick = this.handleAddClick.bind(this);
        // this.handleRemoveClick = this.handleRemoveClick.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleNameChange = e => {
        this.setState({
            [e.target.name]: e.target.value
        });
    }

    handleValueChange = idx => e => {
        const newvalue_list = this.state.value_list.map((value, sidx) => {
            if (idx !== sidx) return value;
            return { ...value, value_name: e.target.value };
        });
        this.setState({ value_list: newvalue_list });
    }
   
    // handle click event of the Remove button
    handleAddClick = () => {
       this.setState({
            value_list: [...this.state.value_list, {value_name: ""}]
        }); 
    };
    
    // handle click event of the Add button
    handleRemoveClick = idx => () => {
        this.setState({
            value_list: this.state.value_list.filter((s, sidx) => idx !== sidx)
        });
    };

    handleSubmit(e) {
        e.preventDefault();
      
        var url = 'http://220.70.2.160:8080/sensor/regist';
        var data = this.state;

        fetch(url, {
        method: 'POST', // or 'PUT'
        body: JSON.stringify(data),
        headers:{
            'Content-Type': 'application/json'
        }
        }).then(res => res.json())
        .then(response => console.log('Success:', JSON.stringify(response)))
        .catch(error => console.error('Error:', error));
    }

    render() {
        return (
            <>
            <button type="button" class="btn btn-primary btn-lg" data-toggle="modal" data-target="#myModal">register sensor</button>
            <div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
                <div class="modal-dialog" role="document">
                    <div class="modal-content">
                        <div class="modal-header"> 
                            <h4 class="modal-title" id="myModalLabel">Register sensor</h4>
                            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">×</span></button>
                        </div>
                        <form>
                            <div class="modal-body">
                                <div class="form-group">
                                    <label for="name">Sensor name</label>
                                    <input type="text" class="form-control" name="name" placeholder="name" value={this.state.name} onChange={this.handleNameChange}/>
                                    <div class="invalid-feedback">
                                        This sensor name is already exist.
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label for="">Value name</label>
                                    {this.state.value_list.map((value_list, idx) => (
                                        <div class="input-group mb-3">
                                            <div class="input-group-prepend">
                                                <span class="input-group-text">{idx}</span>
                                            </div>
                                            <input type="text" class="form-control" name="value_list" placeholder={"Enter value name"} value={value_list.value_name} onChange={this.handleValueChange(idx)}/>
                                            <div class="input-group-append">
                                                <button class="btn btn-primary btn-sm" type="button" id="button-addon2" onClick={this.handleRemoveClick(idx)}>
                                                    <svg width="1em" height="1em" viewBox="0 0 16 16" class="bi bi-trash-fill" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
                                                        <path fill-rule="evenodd" d="M2.5 1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1H3v9a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V4h.5a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1H2.5zm3 4a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 .5-.5zM8 5a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7A.5.5 0 0 1 8 5zm3 .5a.5.5 0 0 0-1 0v7a.5.5 0 0 0 1 0v-7z"/>
                                                    </svg>
                                                </button>
                                            </div>
                                        </div>
                                    ))}
                                </div>
                                <button type="button" class="btn btn-primary" onClick={this.handleAddClick}>Add value</button>
                            </div>
                            <div class="modal-footer">
                                <button type="submit" class="btn btn-primary" data-dismiss="modal" onClick={this.handleSubmit}>Submit</button>
                                <button type="button" class="btn btn-default" data-dismiss="modal">Cancel</button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
            </>
        );
    }
}

export default RegisterSensor;