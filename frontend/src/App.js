import './App.css';
import React, { Component } from 'react';

class Challenge extends Component  {
  constructor() {
    super();
    this.state = {
      count: 1,
      message: '1'
    }
  }
  
  render() { 
    const numberStyle = {
      fontSize: '50px',
    };
    
    return (
      <div  className="box">
        <style>{"body { background-color: #ffa500; }"}</style>
        <h3
          style={numberStyle}>
          {this.state.message}
        </h3>
        <div >
          <button size="large" className="button"
            onClick={this._dec.bind(this)}> 
            - 
          </button>
          <button size="large" className="button"
            onClick={this._inc.bind(this)}>
              +
          </button>
        </div>
      </div>
    );
  }

  fetchAndLog = async () => {
    const response = await fetch('http://localhost:1986/dividend/'+this.state.count);
    const json = await response.json();
    console.log(json);
    this.setState({message: json});
  }
  
  _dec() {
    if (this.state.count > 1) {
      this.setState(
        {
          count: --this.state.count
        }
      );
      this.fetchAndLog();
    }
  }
  
  _inc() {
    if (this.state.count < 100) {
      this.setState(
        {
          count: ++this.state.count
        }
      );
      this.fetchAndLog();
    }
  }
}

export default Challenge;
