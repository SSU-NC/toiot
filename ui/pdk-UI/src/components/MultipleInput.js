import React, { Component } from 'react';

class MultipleInput extends Component {
    /*constructor(props) {
        super(props);
        this.state = { inputs: ['input-0'] };
    }

    render() {
        return(
            <>
               <form>
                   <div class="form-group">
                       {this.state.inputs.map(input => <FormInput key={input} />)}
                   </div>
               </form>
               <button type="button" class="btn btn-default" onClick={ () => this.appendInput() }>
                   {this.props.button_text}
               </button>
            </>
        );
    }

    appendInput() {
        var newInput = `input-${this.state.inputs.length}`;
        this.setState(prevState => ({ inputs: prevState.inputs.concat([newInput]) }));
    }*/
}

export default MultipleInput;