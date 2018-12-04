import React from 'react'
import { Dropdown } from 'semantic-ui-react'

import { friendOptions } from '../common.js';

const DropdownImage = () => (
  <div>
    <Dropdown text='Add user' icon='add user' floating labeled button className='icon'>
      <Dropdown.Menu>
        <Dropdown.Header content='People You Might Know' />
        {/* {friendOptions.map(option => <Dropdown.Item key={option} {...option} />)} */}
      </Dropdown.Menu>
    </Dropdown>
  </div>
)

export default DropdownImage