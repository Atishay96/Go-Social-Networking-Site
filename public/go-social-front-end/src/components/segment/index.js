import React, { Component } from 'react'
import { Segment } from 'semantic-ui-react'

class SegmentReact extends Component{
  constructor() {
    super();
  }
  render() {
    return (
    <div>
      <Segment className="segment_main" raised>
        <div className="question"> Q: { this.props.question } </div>
        <div className="answer"> <b>A: </b>{ this.props.answer } </div>
      </Segment>
    </div> 
    )
  }
}
export default SegmentReact
