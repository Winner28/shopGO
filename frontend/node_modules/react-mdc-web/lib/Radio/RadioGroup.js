'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _List = require('../List/List');

var _List2 = _interopRequireDefault(_List);

var _ListItem = require('../List/ListItem');

var _ListItem2 = _interopRequireDefault(_ListItem);

var _FormField = require('../FormField');

var _FormField2 = _interopRequireDefault(_FormField);

var _Radio = require('./Radio');

var _Radio2 = _interopRequireDefault(_Radio);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var propTypes = {
  className: _propTypes2.default.string,
  children: _propTypes2.default.arrayOf(function (propValue, key, componentName) {
    var type = propValue[key].type;

    if (type !== _Radio2.default) {
      return new Error('Invalid prop \'children[' + key + ']\' supplied to \'' + componentName + '\', expected \'Radio\'');
    }
    return false;
  }),
  name: _propTypes2.default.string.isRequired,
  onChange: _propTypes2.default.func,
  value: _propTypes2.default.oneOfType([_propTypes2.default.number, _propTypes2.default.string])
};

var RadioGroup = function RadioGroup(_ref) {
  var className = _ref.className,
      name = _ref.name,
      onChange = _ref.onChange,
      value = _ref.value,
      otherProps = _objectWithoutProperties(_ref, ['className', 'name', 'onChange', 'value']);

  var childs = _react.Children.map(otherProps.children, function (child) {
    var _child$props = child.props,
        id = _child$props.id,
        children = _child$props.children,
        childValue = _child$props.value;

    var customId = id || 'radio-' + (name || '') + '-' + childValue;
    var radio = (0, _react.cloneElement)(child, {
      checked: childValue === value,
      id: customId,
      name: name,
      onChange: onChange
    });
    return _react2.default.createElement(
      _ListItem2.default,
      null,
      _react2.default.createElement(
        _FormField2.default,
        { id: customId },
        radio,
        _react2.default.createElement(
          'label',
          null,
          children
        )
      )
    );
  });
  return _react2.default.createElement(
    _List2.default,
    { className: className },
    childs
  );
};
RadioGroup.propTypes = propTypes;
exports.default = RadioGroup;