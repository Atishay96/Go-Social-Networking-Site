import React, { Component } from 'react'
import { Segment } from 'semantic-ui-react'

class SegmentReact extends Component {
  render() {
    return (
      <div>
        <Segment className="segment_main" raised>
          <div className="question">{this.props.question} </div>
          <div className="answer"> <b> </b>{this.props.answer} </div>
        </Segment>
      </div>
    )
  }
}
export default SegmentReact
